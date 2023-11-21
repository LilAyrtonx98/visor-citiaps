<template>
  <div v-show="isActive" class="widgetContainer" id="measureID">

      <div class="widgetHeader" id="measureIDheader">
        <p style="font-weight: bold;">Medir mapa</p>
        <div>
          <button class="btn btn-info" v-on:click="isHidden = !isHidden">
            <font-awesome-icon v-if="!isHidden" icon="minus"></font-awesome-icon>
            <font-awesome-icon v-else icon="plus"></font-awesome-icon>
          </button>
          <button class="btn btn-danger" v-on:click="exitMeasure()">
            <font-awesome-icon icon="times"></font-awesome-icon>
          </button>
        </div>
      </div>

      <div class="widgetBody" v-show="!isHidden">
        <label for="mode">Modo</label>
        <br>
        <select id="mode"
          v-if="options"
          v-model="mode"
          v-on:change="selectOption()">
            <option
            v-for="item in options"
            v-bind:key="item.index"
            v-bind:value="item">
            {{item.name}}
            </option>
        </select>

        <!--
        <label for="visibility">Visibilidad de capas</label>
        <div id="visibility">
            <input :disabled="mode.code == 'distance'" type="checkbox" id="checkbox" v-model="visibility.distance" v-on:change="visibilityChange('distance', visibility.distance)">
            <label for="checkbox">Distancia</label>
            <br>
            <input :disabled="mode.code == 'area'" type="checkbox" id="checkbox" v-model="visibility.area" v-on:change="visibilityChange('area', visibility.area)">
            <label for="checkbox">Área</label>
        </div>
        -->

        <button class="btn btn-secondary btn-clean" style="margin-top: 0.5em;" v-on:click="cleanMode()">Limpiar {{mode.name}}</button>
      </div>

  </div>
</template>

<script>
import {draggableDiv} from '@/components/mixin/draggableDiv.js'
import {vuex} from '@/components/mixin/vuex.js'

export default {
  name: 'HelloWorld',
  props: ['isFixed', 'currentMode'],
  mixins:[vuex, draggableDiv],
  data() {
    return {
      options: [
        { name: 'Distancia', code: 'distance'},
        { name: 'Área', code: 'area'},
        { name: 'Radio', code: 'radius'}
      ],
      mode: { name: 'Distancia', code: 'distance'},
      visibility: {
        distance: true,
        area: true
      },
      isHidden: false
    }
  },
  components:{

  },
  computed: {
    isActive: function() {
      return this.isFixed? this.isMobile : this.mapToolValues('measures').active && !this.isMobile
    }
  },
  methods:{
    selectOption: function() {
      // Cambiar interacción en el mapa
      this.executeMapToolAction('measures', 'selectMeasure', {mode: this.mode.code})
      // Cambiar botones de visibilidad en la herramienta
      switch (this.mode.code) {
        case 'distance':
          if(!this.visibility.distance){
            this.visibilityChange(this.mode.code, true);
            this.visibility.distance = true;
          }
          break;
        case 'area':
          if(!this.visibility.area){
            this.visibilityChange(this.mode.code, true);
            this.visibility.area = true;
          }
          break;
        default:
          break;
      }
    },
    cleanMode: function() {
      this.executeMapToolAction('measures', 'clean', {mode: this.mode.code})
    },
    exitMeasure: function() {
      // Limpiar todas las medidas antes de cerrar la herramienta
      this.executeMapToolAction('measures', 'clean', {mode: 'all'})
      // Cerrar la herramienta
      this.selectMapTool('measures', false)
    },
    visibilityChange: function(code, isVisible) {
      this.executeMapToolAction('measures', 'visibilityChange', {mode: code, isVisible: isVisible})
    } 
  },
  updated () {

  },
  created () {

  },
  mounted () {
    this.makeDraggable("measureID");
  },
  watch: {
    // Acción que se ejecuta cuando se activa esta herramienta
    'mapToolSelected': function (newValue, oldValue) {
      if(newValue == 'measures') {
        if(this.isActive) {
          this.executeMapToolAction('measures', 'selectMeasure', {mode: this.mode.code})
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
