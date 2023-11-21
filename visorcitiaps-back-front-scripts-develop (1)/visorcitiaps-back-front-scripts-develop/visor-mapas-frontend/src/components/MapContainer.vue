<template>
<main class="map-page">
  <fullscreen ref="fullscreen" @change="fullscreenChange">

      <LoadScreen
      >
      </LoadScreen>

      <Legends
      ref="legends"
      v-bind:isActive="showMapLegend"
      v-bind:isFixed="false"
      v-on:onLegendAction="legendActionListener"
      >
      </Legends>

      <NavDrawer
      ref="navDrawer"
      >
      </NavDrawer>

      <!--
        v-on:
          onClickFullScreen: captura la acción de apretar el control de activar fullscreen desde el mapa
          onClickMapLegend: captura la acción de apretar el control de mostrar simbologia desde el mapa
      -->

      <MapView
      ref="mapView"
      v-bind:class="{mapcontainermax: fullscreen}"
      v-on:onClickFullScreen="toggleFullScreen()"
      v-on:onClickMapLegend="clickMapLegend"

      v-shortkey="{
        zoomIn: ['shift', 'arrowup'], 
        zoomOut: ['shift', 'arrowdown'],
        openLegend: ['ctrl', 's'],
        getGeolocation: ['ctrl', 'u'],
        fullscreen: ['ctrl', 'p'],
        openDrawer: ['ctrl', 'h'],
        openHelp: ['ctrl', 'a']
        }"
      @shortkey.native="shortkeysHandler"

      >
      </MapView>

      <Select
      v-bind:isFixed="false"
      >       
      </Select>

      <Layers
      v-bind:isFixed="false"
      >
      </Layers>

      <Charts
      v-bind:isFixed="false"
      >
      </Charts>

      <Measure
      v-bind:isFixed="false"
      >
      </Measure>

      <Annotations
      v-bind:isFixed="false"
      >
      </Annotations>

      <Geoprocessings
      v-bind:isFixed="false"
      >
      </Geoprocessings>

      <Search
      v-bind:isFixed="false"
      >
      </Search>

      <Export
      >
      </Export>

      <ChangeBasemap
      >
      </ChangeBasemap>

      <Help
      >
      </Help>

      <AttributesTable
      v-bind:isFixed="false"
      ></AttributesTable>

      <GeoprocessingResult
      >
      </GeoprocessingResult>

  </fullscreen>
</main>
</template>

<style lang="scss" scoped>

body main {
	margin: 0 0%;
	padding: 0% 0;
}

.mapcontainermax{
  background: rgb(3,68,12);  
  height: 100vh;
}
</style>

<script>
// Controles del mapa
import MapView from '@/components/MapView.vue'
import Legends from '@/components/mapControls/legends/Legends.vue'
import LoadScreen from '@/components/mapControls/LoadScreen.vue'
import NavDrawer from '@/components/mapControls/NavDrawer.vue'
import Fullscreen from "vue-fullscreen/src/component.vue"

// Herramientas del mapa
import Select from '@/components/mapTools/select/Select.vue'
import Layers from '@/components/mapTools/layers/Layers.vue'
import Charts from '@/components/mapTools/charts/Charts.vue'
import Measure from '@/components/mapTools/measures/Measure.vue'
import Geoprocessings from '@/components/mapTools/geoprocessings/Geoprocessings.vue'
import Search from '@/components/mapTools/search/Search.vue'
import Export from '@/components/mapTools/export/Export.vue'
import ChangeBasemap from '@/components/mapTools/baseMap/ChangeBasemap.vue'
import Help from '@/components/mapTools/help/Help.vue'
import GeoprocessingResult from '@/components/mapTools/geoprocessings/GeoprocessingResult.vue'
import Annotations from '@/components/mapTools/annotations/Annotations.vue'
import AttributesTable from '@/components/mapTools/attributesTable/AttributesTable.vue'

// Mixin
import {restMap} from '@/components/mixin/restMap.js'
import {restVisor} from '@/components/mixin/restVisor.js'
import {vuex} from '@/components/mixin/vuex.js'

export default{
  name: 'home',
  mixins: [restVisor, restMap, vuex],
  data(){
    return{
      fullscreen: false,
      showMapLegend: false,
      }
  },
  components:{
      MapView,
      NavDrawer,
      Fullscreen,
      Measure,
      Select,
      Layers,
      Charts,
      Search,
      ChangeBasemap,
      LoadScreen,
      Legends,
      Help,
      Geoprocessings,
      Export,
      GeoprocessingResult,
      Annotations,
      AttributesTable
  },
  computed: {
  },
  props: [],
  methods:{
    shortkeysHandler: function(event) {
      switch (event.srcKey) {
        case 'openLegend':
          this.clickMapLegend();
          break;
        case 'getGeolocation':
          this.$refs.mapView.getGeolocation();
          break;
        case 'fullscreen':
          this.toggleFullScreen();
          break;
        case 'openDrawer':
          this.selectMapControl('navDrawer', true)
          break;
        case 'openHelp':
         this.selectMapTool('help', true)
          break;
        default:
          this.$refs.mapView.shortkeysHandler(event.srcKey);
          break;
      }
    },
    toggleFullScreen () {
      this.$refs['fullscreen'].toggle()
    },
    fullscreenChange (fullscreen) {
      this.fullscreen = fullscreen
    },
    legendActionListener: function(widget, action, object) {
      switch (widget) {
        case 'legend':
          if(action == 'exit') {
            this.clickMapLegend()
          }
          break;
        default:
          break;
      }
    },
    clickMapLegend: function() {
      if(this.$store.getters.isMobile) {
        this.$refs.navDrawer.openNavAnimation();
        this.$refs.navDrawer.showMapLegend();
      } else {
        this.showMapLegend = !this.showMapLegend;
      }
    }
  },
  updated () {
    
  },
  created () {

  },
  mounted () {
    // Mostrar footer
    this.setMapActive(true)
    //this.cleanActiveLayers()
  },
  beforeDestroy() {
    // Limpiar el estado del mapa antes de salir
    this.resetMap()
    // Ocultar footer
    this.setMapActive(false)
  }
  
}
</script>
