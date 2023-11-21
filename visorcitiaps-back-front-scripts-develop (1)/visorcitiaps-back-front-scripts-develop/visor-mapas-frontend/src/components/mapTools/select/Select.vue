<template>
  <div v-show="isActive" :class="{widgetContainer : !isFixed}" :id="isFixed? 'selectIDFix' : 'selectID'">
    <div v-if="!isFixed" :class="{widgetHeader : !isFixed}" id="selectIDheader">
        <p style="font-weight: bold;">Selección</p>
        <div>
          <button class="btn btn-info" v-on:click="isHidden = !isHidden">
            <font-awesome-icon v-if="!isHidden" icon="minus"></font-awesome-icon>
            <font-awesome-icon v-else icon="plus"></font-awesome-icon>
          </button>
          <button class="btn btn-danger" v-on:click="exitSelect()">
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
      mode: { name: 'Seleccionar punto', code: 'point', drawType: 'Point'},
      isHidden: false
    }
  },
  components:{

  },
  computed: {
    isActive: function() {
      // Si isFixed = true, componente está activo si:
      //  - es un dispositivo móvil
      // Si no, componente está activo si:
      //  - la herramienta 'select' está seleccionada
      //  - no es un dispositivo móvil
      return this.isFixed? this.isMobile : this.mapToolValues('select').active && !this.isMobile
    },
    options: function() {
      let optionsArray = [
        { name: 'Seleccionar punto', code: 'point', drawType: 'Point'}
      ]
      this.isMobile? true : optionsArray.push({ name: 'Seleccionar polígono', code: 'area', drawType: 'Polygon'}), 
                            optionsArray.push({ name: 'Seleccionar círculo', code: 'area', drawType: 'Circle'})
      return optionsArray
    }
  },
  methods:{
    selectOption: function() {
      // Cambiar interacción en el mapa
      this.executeMapToolAction('select', 'selectMode', {mode: this.mode.code, drawType: this.mode.drawType})
    },
    exitSelect: function() {
      // Cerrar la herramienta
      this.selectMapTool('select', false)
    }
  },
  updated () {

  },
  created () {

  },
  mounted () {
    this.makeDraggable("selectID");
  },
  watch: {
    // Acción que se ejecuta cuando se activa esta herramienta
    'mapToolSelected': function (newValue, oldValue) {
      if(newValue == 'select') {
        if(this.isActive) {
          this.executeMapToolAction('select', 'selectMode', {mode: this.mode.code, drawType: this.mode.drawType})
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
