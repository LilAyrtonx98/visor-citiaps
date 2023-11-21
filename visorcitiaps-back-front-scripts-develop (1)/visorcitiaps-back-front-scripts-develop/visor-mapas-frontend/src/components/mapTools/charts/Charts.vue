<template>
  <div v-show="isActive" :class="{widgetContainer : !isFixed}" :id="isFixed? 'chartsIDFix' : 'chartsID'">
    <div v-if="!isFixed" :class="{widgetHeader : !isFixed}" id="chartsIDheader">
        <p>Gráficos Estadísticos</p>
        <div>
          <button class="btn btn-info" v-on:click="isHidden = !isHidden">
            <font-awesome-icon v-if="!isHidden" icon="minus"></font-awesome-icon>
            <font-awesome-icon v-else icon="plus"></font-awesome-icon>
          </button>
          <button class="btn btn-danger" v-on:click="exitCharts()">
            <font-awesome-icon icon="times"></font-awesome-icon>
          </button>
        </div>
    </div>

    <div :class="{widgetBody : !isFixed}" v-show="!isHidden">
      <label for="chart-data-select-1" style="font-size: 1em;">Seleccione una capa</label>
      <br>
      <select id="chart-data-select-1" style="margin-bottom: 0.5em;" v-model="chartLayerType">
          <option value="">
            Ninguna capa
          </option>
          <option
          value="chileRegiones2018">
            Regiones
          </option>
          <option
          value="chileComunas2018">
            Comunas
          </option>
      </select>
      <div class="btn-container-no-shadow">
        <button :disabled="chartLayerType =='' ? true:false" class="btn btn-secondary" style="border-radius: 6px;" :class="{btnSelected : chart==0}" v-on:click="displayDataInChart(0)">
          <span>Población</span>
        </button>
        <button :disabled="chartLayerType =='' ? true:false" class="btn btn-secondary" style="border-radius: 6px;" :class="{btnSelected : chart==1}" v-on:click="displayDataInChart(1)">
          <span>Viviendas</span>
        </button>
        <button :disabled="chartLayerType =='' ? true:false" class="btn btn-secondary" style="border-radius: 6px;" :class="{btnSelected : chart==2}" v-on:click="displayDataInChart(2)">
          <span>Incendios</span>
        </button>
      </div>

      <div :class="{resizable : !isFixed}">
        <HistogramChart
        v-if="chart == 0 && chartLayerType !=''"
        v-bind:chartData="populationData"
        v-bind:showAltView="true"
        ref="histogramChart">
        </HistogramChart>

        <PieChart2
        v-if="chart == 1 && chartLayerType !=''"
        v-bind:chartData="houseData"
        v-bind:showAltView="true"
        ref="pieChart2">
        </PieChart2>

        <AreaChart
        v-if="chart == 2 && chartLayerType !=''"
        v-bind:chartData="fireData"
        v-bind:showAltView="true"
        ref="areaChart">
        </AreaChart>

        <div v-if="chartLayerType ==''">
          <ChartStatNoData
          v-bind:chartContainerAlt="true"
          v-bind:type="1"/>
        </div>
      </div>
      
    </div>

  </div>
</template>

<script>
import {draggableDiv} from '@/components/mixin/draggableDiv.js'
import {vuex} from '@/components/mixin/vuex.js'

import HistogramChart from '@/components/mapTools/charts/chartType/HistogramChart.vue';
import AreaChart from '@/components/mapTools/charts/chartType/AreaChart.vue';
import PieChart2 from '@/components/mapTools/charts/chartType/PieChart2.vue';
import ChartStatNoData from '@/components/mapTools/charts/ChartStatNoData.vue';

export default {
  name: 'HelloWorld',
  props: ['isFixed'],
  mixins:[draggableDiv, vuex],
  data() {
    return {
      chart: 0,
      chartLayerType: '',
      populationData: {isData: false, series: [0], chartOptions: {}},
      fireData: {isData: false, series: [0], chartOptions: {}},
      houseData: {isData: false, series: [0], chartOptions: {}},

      isHidden: false,
    }
  },
  components:{
    HistogramChart,
    AreaChart,
    PieChart2,
    ChartStatNoData
  },
  computed: {
    isActive: function() {
      return this.isFixed? this.isMobile : this.mapToolValues('chart').active && !this.isMobile
    }
  },
  methods:{
    onShowDataChart: function(chartData) {
      this.populationData = chartData[0];
      this.fireData = chartData[1];
      this.houseData = chartData[2];
    },
    displayDataInChart: function(chart) {
      this.chart = chart;
    },
    exitCharts: function() {
      // Quitar datos del gráfico
      this.populationData = {isData: false, series: [0], chartOptions: {}};
      this.fireData = {isData: false, series: [0], chartOptions: {}};
      this.houseData = {isData: false, series: [0], chartOptions: {}};
      // Cerrar widget
      this.selectMapTool('chart', false)
    },
  },
  updated () {

  },
  created () {

  },
  mounted () {
    this.makeDraggable("chartsID");
  },
  watch: {
    'chartLayerType': function (newValue, oldValue) {
     this.executeMapToolAction('chart', 'addLayer', {layer: newValue})

      if(newValue == '') {
        this.populationData = {isData: false, series: [0], chartOptions: {}};
        this.fireData = {isData: false, series: [0], chartOptions: {}};
        this.houseData = {isData: false, series: [0], chartOptions: {}};
      }
    },
    'mapToolExecuted': function (newValue, oldValue) {
      switch (newValue) {
        case 'chart':
          if(this.mapToolValues('chart').action == 'showDataChart') {
            this.onShowDataChart(this.mapToolValues('chart').data.chartData)
          }
          break;
      
        default:
          break;
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

.resizable{
  margin-top: 0.5em;
  min-height: 20vh;
  min-width: 100%;
  resize: both;
  overflow: auto;
}

.btn-clean {
  border-radius: 12px;
  -webkit-transition: 200ms ease;
  font-size: 0.75em;
  transition: 200ms ease;
  font-weight: bold;
  padding: 0.3em 0.9em;
  min-width: 8em;
}

.btnSelected {
	color: white;
  background: #036016;
}
.btnTabSelected {
	color: white;
  background: #036016;
  box-shadow: -2px -2px 0.25em gray;
}
.btn-tab{
  border-bottom: solid 2px gray;
}

.btn-container-no-shadow {
  display: flex;
  justify-content: space-evenly;
  .btn-secondary{
    justify-content: center;
    border: solid 2px #03440c;
    width: 2em;
    height: 2em;
  }
  .btn-secondary:hover{
    color: #03440c;
    background: #f7e200;
  }
  .btn-secondary:disabled {
    color: gray;
    background: white;
    border: solid 2px gray;
  }

  .btn-secondary:disabled:hover {
    background: none;
    color: gray;
    opacity: 1;
  }
}

</style>
