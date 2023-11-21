<template>
    <modal
        name="GeoprocessingResult"
        style="z-index: 100000001"
        :height="'auto'"
        :width="'100%'"
    >
        <div class="modal-content">
            <div class="modal-item container-1">
                <h2>Resultado geoproceso</h2>
                
                <div v-if="barData.isData">
                    <label
                    for="layer-select"
                    >
                    Capa
                    </label>
                    <select id="layer-select"
                    v-model="selectedLayer"
                    >
                        <option
                        v-for="item in layersResult"
                        v-bind:key="item.index"
                        v-bind:value="item">
                        {{item}}
                        </option>
                    </select>

                    <div
                    style="
                    display: flex;
                    flex-direction: column;
                    align-items: start;
                    margin-top: 0.5em;"
                    >
                        <strong>{{selectedPage + '/' + barData.series.length}}</strong>

                        <h2>
                        {{barData.series[selectedSerie].name}} 
                        </h2>
                    </div>
                </div>

                <BarChart v-bind:chartData="barDataOneSerie(this.selectedSerie)" v-bind:showAltView="true" ref="barChart"></BarChart>
            </div>

            <div
            style="
            display: flex;
            justify-content: space-between;"
            >
                <button v-on:click="hideModal()" class="btn btn-secondary">Volver</button>
                <div v-if="barData.isData" class="buttons-container modal-buttons">
                    <button v-on:click="changeSerie(-1)" class="btn btn-secondary" style="min-width: 4em;">
                        <font-awesome-icon icon="chevron-left"/>
                    </button>
                    <button v-on:click="changeSerie(1)" class="btn btn-secondary" style="min-width: 4em;">
                        <font-awesome-icon icon="chevron-right"/>
                    </button>
                </div>
            </div>

            <div class="buttons-container modal-buttons">
                
            </div>
        </div>
    </modal>
</template>

<script>
import {formatter} from "@/components/mixin/formatter.js";
import {geoprocessingResult} from "@/components/mapTools/geoprocessings/geoprocessingResult.js";
import {vuex} from '@/components/mixin/vuex.js'

import BarChart from "@/components/mapTools/charts/chartType/BarChart.vue";

export default {
    mixins: [formatter, geoprocessingResult, vuex],
    props: [],
    components: {
        BarChart
    },
    data() {
        return {
            barData: { isData: false, series: [0], chartOptions: {} },
            selectedSerie: 0,
            selectedLayer: '',
        };
    },
    computed: {
        selectedPage: function() {
            return this.selectedSerie + 1
        }
    },
    methods: {
        showModal: function() {
            this.$modal.show("GeoprocessingResult");
            this.selectedLayer == '' ? this.selectedLayer = this.layersResult[0]: false
            if(this.selectedLayer == '' && this.selectedLayer!=undefined) {
                this.onShowDataChart(this.selectedLayer)
            }
        },
        hideModal: function() {
            this.$modal.hide("GeoprocessingResult");
        },
        onShowDataChart: function(chartData) {
            this.barData = this.transformDataForChart(this.dataChart[chartData])
        },
        barDataOneSerie: function(serie) {
            let clone = Object.assign({}, this.barData)
            clone.series = []
            clone.series.push(this.barData.series[serie])
            return clone
        },
        changeSerie: function(num) {
            if(this.selectedSerie + num >= 0 && this.selectedSerie + num < this.barData.series.length) {
                this.selectedSerie = this.selectedSerie + num
            }
        }
    },
    created() {},
    updated() {},
    watch: {
        'selectedLayer': function (newValue, oldValue) {
            this.selectedSerie = 0;
            this.onShowDataChart(this.selectedLayer)
        },
        'mapToolExecuted': function (newValue, oldValue) {
            switch (newValue) {
                case 'geoprocessing':
                    if(this.mapToolValues('geoprocessing').action == 'showResults') {
                        this.showModal()
                    } else if(this.mapToolValues('geoprocessing').action == 'sendAttributes') {
                        this.receiveAttributes(this.mapToolValues('geoprocessing').data.attributes)
                    }
                    break;
                default:
                    break;
            }
        }
    },
    mounted() {

    }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
p {
    margin-left: 1em;
}
.button-explanation {
    display: flex;
    flex-direction: row;
    width: 100%;
    align-items: center;
}
.key {
    background-color: gray;
    width: 2em;
    height: 2em;
    text-align: center;
    color: white;
    border-radius: 4px;
    font-size: 1.25em;
    padding: 0.25em;
    font-weight: bold;
}
</style>
