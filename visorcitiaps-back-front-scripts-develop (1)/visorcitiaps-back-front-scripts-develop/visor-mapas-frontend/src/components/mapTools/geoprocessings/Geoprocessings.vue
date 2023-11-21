<template>
  <div v-show="isActive" class="widgetContainer" id="geoprocessingID">

      <div class="widgetHeader" id="geoprocessingIDheader">
        <p style="font-weight: bold;">Ejecutar Geoproceso</p>
        <div>
          <button class="btn btn-info" v-on:click="isHidden = !isHidden">
            <font-awesome-icon v-if="!isHidden" icon="minus"></font-awesome-icon>
            <font-awesome-icon v-else icon="plus"></font-awesome-icon>
          </button>
          <button class="btn btn-danger" v-on:click="exitGeoprocessing()">
            <font-awesome-icon icon="times"></font-awesome-icon>
          </button>
        </div>
      </div>

      <div class="widgetBody" v-show="!isHidden">
        <label for="geo">Geoproceso</label>
        <br>
        <select id="geo"
          v-if="geoprocessings"
          v-model="geoprocessingUrl"
          >
            <option
            value="-1"
            disabled
            >
            Seleccionar...
            </option>
            <option
            v-for="item in geoprocessings"
            v-bind:key="item.index"
            v-bind:value="item.geo_url">
            {{item.name}}
            </option>
        </select>



        <label for="draw">Modo de selección</label>
        <br>
        <select id="draw"
          v-model="drawSelected"
          v-on:change="drawChange()"
          >
            <option
            v-for="item in draws"
            v-bind:key="item.index"
            v-bind:value="item">
            {{item.name}}
            </option>
        </select>


        <div id="results" style="color: red;">
            <span v-if="activeLayers.length<=0">No hay capas en el mapa</span>
            <span v-else-if="geoprocessingUrl=='-1'">No se ha elegido un geoproceso</span>
            <span v-else-if="errorMessage != 'ok'">{{errorMessage}}</span>
            <br>
            <div style="display: flex;">
              <button class="btn btn-secondary btn-clean" style="margin-top: 0.5em; width: 50%; white-space: nowrap;" v-on:click="cleanResults()">Borrar Resultados</button>
              <button :disabled="activeLayers.length<=0 || geoprocessingUrl=='-1' || errorMessage!='ok'" class="btn btn-secondary btn-clean" style="margin-top: 0.5em; width: 50%; white-space: nowrap;" v-on:click="showResults()">Ver Resultados</button>
            </div>
        </div>

        <div style="display: flex;">
          <button class="btn btn-secondary btn-clean" style="margin-top: 0.5em; width: 50%; white-space: nowrap;" v-on:click="cleanMode()">Borrar Polígono</button>
          <button :disabled="activeLayers.length<=0 || geoprocessingUrl=='-1'" class="btn btn-secondary btn-clean" style="margin-top: 0.5em; width: 50%; white-space: nowrap;" v-on:click="executeGeoprocessing()" >Ejecutar</button>
        </div>
      </div>

  </div>
</template>

<script>
import {draggableDiv} from '@/components/mixin/draggableDiv.js'
import {vuex} from '@/components/mixin/vuex.js'

export default {
  name: 'HelloWorld',
  props: ['isFixed'],
  mixins:[draggableDiv, vuex],
  data() {
    return {
      errorMessage: '',

      geoprocessingBody: {
        layers: [], 
        geometry: {
          coordinates: [],
          type: "Polygon"
        },
        radius: 1
      },

      geoprocessingUrl: "-1",
      drawSelected: {name: 'Polígono', code: 'Polygon'},
      draws: [
        {name: 'Polígono', code: 'Polygon'},
        {name: 'Círculo', code: 'Circle'}
      ],

      isHidden: false
    }
  },
  components:{

  },
  computed: {
    isActive: function() {
      return this.isFixed? this.isMobile : this.mapToolValues('geoprocessing').active && !this.isMobile
    }
  },
  methods:{
    cleanMode: function() {
      this.executeMapToolAction('geoprocessing', 'clean', {})
    },
    exitGeoprocessing: function() {
      this.cleanResults()
      this.selectMapTool('geoprocessing', false)
      this.errorMessage = ''
    },
    executeGeoprocessing: function() {

      this.geoprocessingBody.layers = []
      
      // Nombre de las capas presentes en el mapa
      
      this.activeLayers.forEach(layer => {
        this.geoprocessingBody.layers.push(layer.provider.geoserverdata.workspace + "." + layer.provider.geoserverdata.filename)
      });

      this.geoprocessingBody.geometry.type = this.drawSelected.code

      this.executeMapToolAction('geoprocessing', 'execute', {
        geoprocessingBody: this.geoprocessingBody,
        geoprocessingUrl: this.geoprocessingUrl
        })
    
    },
    showResults: function() {
      this.executeMapToolAction('geoprocessing', 'showResults', {})
    },
    cleanResults: function() {
      this.errorMessage = ''
      this.executeMapToolAction('geoprocessing', 'cleanResults', {})
    },
    showError: function(error) {
      this.errorMessage = error
    },
    drawChange: function() {
      this.executeMapToolAction('geoprocessing', 'changeDrawType', {drawType: this.drawSelected})
    } 
  },
  updated () {

  },
  created () {

  },
  mounted () {
    //
    this.makeDraggable("geoprocessingID");
  },
    watch: {
    'mapToolExecuted': function (newValue, oldValue) {
      switch (newValue) {
        case 'geoprocessing':
          if(this.mapToolValues('geoprocessing').action == 'sendError') {
            this.showError(this.mapToolValues('geoprocessing').data.errorMessage)
          }
          break;
      
        default:
          break;
      }
    },
    'mapToolSelected': function (newValue, oldValue) {
      if(newValue == 'geoprocessing') {
        if(this.isActive){
          this.executeMapToolAction('geoprocessing', 'changeDrawType', {drawType: this.drawSelected})
        }
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

.measure-container {
  position: absolute;
  left: 0.5em;
  top: 4em;
  padding: 0.25em;
  
  background: rgba(255, 255, 255, 0.5);
  box-shadow: 2px 2px 1px grey;
  border-radius: 4px;
  border: solid 2px rgba(3, 96, 22, 0.5);
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

.btn-danger {
  font-size: 0.75em;
  min-width: 2em;
}

</style>
