<template>

<div id="mySidenav" class="sidenav">

  <div style="display: flex; align-items: center; justify-content: space-between; background: #f1f1f1; height: 4em;">
    <strong style="margin-left: 1em;">{{title}}</strong>
    <a style="margin-right: 1em;" href="javascript:void(0)" class="closebtn" v-on:click="closeNav()">
      <font-awesome-icon icon="chevron-left" style="font-size: 2em;"></font-awesome-icon>
    </a>
  </div>

  <div class="btn-container tab-container" v-show="isMobile&&mode!=4">
    <button class="btn btn-secondary btn-tab" v-on:click="mode = 0; resetAllSelectedButtons()" :class="{btnTabSelected : mode==0}">
      <font-awesome-icon icon="arrow-left" v-if="mode!=0"></font-awesome-icon>
      <font-awesome-icon icon="toolbox" v-else></font-awesome-icon>
    </button>
    <button class="btn btn-secondary btn-tab" v-show="mode==1" v-on:click="mode = 1" :class="{btnTabSelected : mode==1}">
      <font-awesome-icon icon="layer-group"></font-awesome-icon>
    </button>
    <button class="btn btn-secondary btn-tab" v-show="mode==2" v-on:click="mode = 2" :class="{btnTabSelected : mode==2}">
      <font-awesome-icon icon="chart-bar"></font-awesome-icon>
    </button>
    <button class="btn btn-secondary btn-tab" v-show="mode==3" v-on:click="mode = 3" :class="{btnTabSelected : mode==3}">
      <font-awesome-icon icon="search"></font-awesome-icon>
    </button>
    <button class="btn btn-secondary btn-tab" v-show="mode==4" v-on:click="mode = 4" :class="{btnTabSelected : mode==4}">
      <font-awesome-icon icon="map"></font-awesome-icon>
    </button>
    <button class="btn btn-secondary btn-tab" v-show="mode==5" v-on:click="mode = 5" :class="{btnTabSelected : mode==5}">
      <font-awesome-icon icon="map"></font-awesome-icon>
    </button>
  </div>

  <div v-show="mode==0" class="tools">
    <!--
    <button class="btn btn-secondary" :class="{btnSelected : activeButtons.select}" v-on:click="activeButtons.select = !activeButtons.select; onClickSelectTool('select', activeButtons.select);">
    -->
    <button class="btn btn-secondary" :class="{btnSelected : mapToolSelected == 'select' && !isMobile}" v-on:click="isMobile? mode=5 : false, onClickSelectTool('select'); isMobile? closeNav():false">
      <font-awesome-layers class="fa-lg">
        <font-awesome-icon icon="hand-pointer"/>
      </font-awesome-layers>
      <p>Selección</p>
    </button>
    <button class="btn btn-secondary" :class="{btnSelected : mapToolSelected == 'layers' && !isMobile}" v-on:click="isMobile? mode=1 : false, onClickSelectTool('layers')">
      <font-awesome-layers class="fa-lg">
        <font-awesome-icon icon="layer-group"/>
      </font-awesome-layers>
      <p>Capas</p>
    </button>
    <button class="btn btn-secondary" :class="{btnSelected : mapToolSelected == 'chart' && !isMobile}" v-on:click="isMobile? mode=2 : false,  onClickSelectTool('chart')">
      <font-awesome-layers class="fa-lg">
        <font-awesome-icon icon="chart-bar"/>
      </font-awesome-layers>
      <p>Gráficos Estadísticos</p>
    </button>
    <button class="btn btn-secondary" v-show="!isMobile" :class="{btnSelected : mapToolSelected == 'measures'}" v-on:click="isMobile? mode=0 : onClickSelectTool('measures')">
      <font-awesome-layers class="fa-lg">
        <font-awesome-icon icon="ruler-combined"/>
      </font-awesome-layers>
      <p>Medir Mapa</p>
    </button>
    <button class="btn btn-secondary" v-show="!isMobile" :class="{btnSelected : mapToolSelected == 'geoprocessing'}" v-on:click="isMobile? mode=0 : onClickSelectTool('geoprocessing')">
      <font-awesome-layers class="fa-lg">
        <font-awesome-icon icon="cogs"/>
      </font-awesome-layers>
      <p>Geoprocesos</p>
    </button>
    <button class="btn btn-secondary" :class="{btnSelected : mapToolSelected == 'search' && !isMobile}" v-on:click="isMobile? mode=3 : false, onClickSelectTool('search')">
      <font-awesome-layers class="fa-lg">
        <font-awesome-icon icon="search"/>
      </font-awesome-layers>
      <p>Buscar Ubicación</p>
    </button>
    <button class="btn btn-secondary" v-on:click="onClickSelectTool('export')">
      <font-awesome-layers class="fa-lg">
        <font-awesome-icon icon="file-export"/>
      </font-awesome-layers>
      <p>Exportar Mapa</p>
    </button>
    <button class="btn btn-secondary" v-on:click="onClickSelectTool('basemap')">
      <font-awesome-layers class="fa-lg">
        <font-awesome-icon icon="map"/>
      </font-awesome-layers>
      <p>Mapa Base</p>
    </button>
    <button class="btn btn-secondary" :class="{btnSelected : mapToolSelected == 'annotations' && !isMobile}" v-on:click="isMobile? mode=1 : false, onClickSelectTool('annotations')">
      <font-awesome-layers class="fa-lg">
        <font-awesome-icon icon="sticky-note"/>
      </font-awesome-layers>
      <p>Anotaciones</p>
    </button>
    <button class="btn btn-secondary" v-on:click="onClickMapHelp()">
      <font-awesome-layers class="fa-lg">
        <font-awesome-icon icon="question"/>
      </font-awesome-layers>
      <p>Ayuda</p>
    </button>
    <button class="btn btn-secondary" v-on:click="onClickGoHome()">
      <font-awesome-layers class="fa-lg">
        <font-awesome-icon icon="times"/>
      </font-awesome-layers>
      <p>Salir</p>
    </button>

  </div>

  <div v-show="mode==1">
    <Layers
    v-bind:isFixed="true"
    >
    </Layers>
  </div>

  <div v-show="mode==2">
    <Charts
    ref="chartsDiv"
    v-bind:isFixed="true"
    >
    </Charts>
  </div>

  <div v-show="mode==3">
      <Search
      v-bind:isFixed="true"
      >
      </Search>
  </div>

  <div v-show="mode==4">
      
    <Legends
    ref="legends"
    v-bind:isFixed="true"
    >
    </Legends>

  </div>

  <div v-show="mode==5">
    <Select
    v-bind:isFixed="true"
    >       
    </Select>
  </div>

</div>

</template>

<script>
import Layers from '@/components/mapTools/layers/Layers.vue'
import Charts from '@/components/mapTools/charts/Charts.vue'
import Search from '@/components/mapTools/search/Search.vue'
import Legends from '@/components/mapControls/legends/Legends.vue';
import Select from '@/components/mapTools/select/Select.vue'
import Annotations from '@/components/mapTools/annotations/Annotations.vue';

import {vuex} from '@/components/mixin/vuex.js'

export default{
  props:[],
  mixins: [
    vuex
  ],
  components: {
    Layers,
    Charts,
    Search,
    Legends,
    Annotations,
    Select
  },
  data(){
    return {
      mode: 0,
      chart: 0,
      activeButtons: {
        select: false,
        measures: false,
        chart: false,
        magnify: false,
        export: false,
        upload: false,
        layers: false,
        search: false,
        geoprocessing: false,
        annotations: false,
      },
      lastMode: 0
    }
  },
  computed: {
    title: function() {
      switch (this.mode) {
        case 0:
          return 'Herramientas'
        case 1:
          return 'Lista de capas'
        case 2:
          return 'Gráficos'
        case 3:
          return 'Buscador'
        case 4:
          return 'Simbología'
        case 5:
          return 'Selección'
        default:
          return 'Listado';
      }
    },
    isActive: function() {
      return this.mapControlValues('navDrawer').active
    }
  },
  methods: {
    openNavAnimation: function() {
      if(this.mode == 4) {
        this.mode = this.lastMode
      }
     if(this.isMobile) {
       document.getElementById("mySidenav").style.width = "100vw";
     } else {
       document.getElementById("mySidenav").style.width = "15em";
     }
    },
    showMapLegend: function() {
      this.lastMode = this.mode;
      this.mode = 4;
    },
    closeNavAnimation: function() {
        document.getElementById("mySidenav").style.width = "0";
    },
    closeNav: function() {
      this.selectMapControl('navDrawer', false)
      this.closeNavAnimation();
    },
    resetAllSelectedButtons: function() {
      this.$store.commit('resetMapToolsSelection')
    },
    onClickSelectTool: function(tool) {
      // Cerrar nav drawer
      this.isMobile? true : this.closeNav()
      // Cambiar herramienta seleccionada
     this.selectMapTool(tool, true)
    },
    onClickGoHome: function() {
      this.selectMapControl('navDrawer', false)
      window.history.length > 1
          ? this.$router.go(-1)
          : this.$router.push('/inicio/')
    },
    onClickMapHelp: function() {
      //this.$router.push('/mapa/ayuda')
      let routeData = this.$router.resolve({name: 'ayuda del mapa'});
      window.open(routeData.href, '_blank');
    }
  },
  created(){
  
  },
  updated() {
      
  },
  watch: {
    'isActive': function(newValue, oldValue) {
        if(newValue){
            this.openNavAnimation();
        } else {
            this.closeNavAnimation();
        }
    }
  }
}
</script>

<style lang="scss" scoped>

.sidenav {
  height: 100%;
  width: 0;
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  background-color: #ffffff;
  overflow-x: hidden;
  transition: 0.5s;
  z-index: 100000000;
}

.sidenav a {
  /*
  padding: 8px 8px 8px 32px;
  text-decoration: none;
  font-size: 25px;
  color: #818181;
  display: block;
  transition: 0.3s;
  */
  text-decoration: none;
  color: #036016;
  transition: 0.3s;
}

.sidenav a:hover {
  color: #f7e200;
}

.sidenav .closebtn {
  /*
  position: absorelute;
  top: 0;
  right: 25px;
  font-size: 36px;
  margin-left: 50px;
  */
}

.btn-container {
  display: flex;
  justify-content: space-evenly;
  .btn-secondary{
    justify-content: center;
    width: 5em;
    height: 2em;
  }
  .btn-secondary:hover{
    color: #03440c;
    background: #f7e200;
    box-shadow: -2px -2px 0.25em gray;
  }
}
.btn-container-no-shadow {
  display: flex;
  flex-direction: column; 
  justify-content: space-evenly;
  align-items: center;
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
.tab-container{
  background: #f1f1f1;
  width: 100%;
  justify-content: flex-start;
  button {
    min-width: 1em;
  }
}
.btn {
  //padding-top: 0.75em;
}
.tools{
  display: flex;
  flex-direction: column;
  grid-column-gap: 0.25em;
  //grid-row-gap: 0.25em;
}
.btn-secondary {
  border: none;
  display: flex;
  border-radius: 0px;
  font-size: 0.9em;
  align-items: center;
  height: 4em;
  background: transparent;
  p {
    margin-left: 0.5em;
  }
}
.toolSelected {
	color: #03440c;
	background: black;
}
p {
  margin-top: 0.5em;
}
.fa-lg {
  font-size: 2em;
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
  width: min-content;
}
.btn-secondary:hover{
  color: #03440c;
  background: #f7e200;
}

@media screen and (max-width: 960px) {
  .tools{
    grid-template-columns: 1fr 1fr;
  }
}

.place-title {
    display: flex;
    align-items: center;
    font-weight: bold;
    font-size: 1.25em;
    img {
      width: 1.25em;
      height: 1.25em;
    }
  }
  .place-name {
    //font-size: 1.0625em;
    //color: #036016;
  }
  .place-lat-lon{
    color: #036016;
  }
  span{
    //font-weight: bold;
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

</style>