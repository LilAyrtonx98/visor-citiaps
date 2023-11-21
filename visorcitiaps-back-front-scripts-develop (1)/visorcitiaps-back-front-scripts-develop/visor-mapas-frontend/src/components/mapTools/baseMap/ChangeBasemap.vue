<template>
  <modal name="ChangeBasemap" style="z-index: 100000001" :height="'auto'" :width="'90%'">
      <div class="modal-content">
          <div class="modal-item container-1">                
              <h2>Selecciona mapa</h2>

            <div class="basemap-style">
                
                <input type="radio" id="openStreetMap" :value="'openStreetMap'" :checked="true" v-model="basemapSelected">
                <label for="openStreetMap">OpenStreetMap</label>
                <br>
                <input type="radio" id="local" :value="'local'" v-model="basemapSelected">
                <label for="local">Local</label>
                <br>

                <div class="mapDescription" v-if="basemapSelected=='openStreetMap'">
                    <p class="name">Mapa de OpenStreetMap</p>
                    <p>Las imágenes provienen de OpenStreetMap.org.</p>
                    <p>Requiere conexión a internet. Es posible visualizar Chile y al resto del mundo.</p>
                </div>

                <div class="mapDescription" v-if="basemapSelected=='local'">
                    <p class="name">Mapa Local</p>
                    <p>Las imágenes están almacenadas localmente.</p>
                    <p>No requiere conexión a internet. Es posible visualizar sólo a Chile.</p>
                </div>
               
            </div>
              
          </div>

          <div class="buttons-container modal-buttons" >
              <button v-on:click="hideModal()" class="btn btn-secondary">Cancelar</button>
              <button v-on:click="onClickOK()" class="btn btn-primary">Confirmar</button>
          </div>

      </div>
  </modal>
</template>

<script>
import {formatter} from '@/components/mixin/formatter.js'
import {vuex} from '@/components/mixin/vuex.js'

export default{
  mixins: [vuex, formatter],
  props:[],
  components: {
    
  },
  data(){
    return {
      basemapSelected: 'openStreetMap',
    }
  },
  computed: {
      isActive: function() {
          return this.mapToolValues('basemap').active
      }
  },
  methods: {
    showModal: function(index){
        this.$modal.show("ChangeBasemap")
    },
    hideModal: function(){
        this.$modal.hide("ChangeBasemap")
    },
    onClickOK: function() {
        this.executeMapToolAction('basemap', 'change', {mode: this.basemapSelected})
    },
  },
  created(){
    
  },
  updated() {
    
  },
  watch: {
    'isActive': function (newValue, oldValue) {
        if(newValue){
            this.showModal()
        } else {
            this.hideModal()
        }
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

.basemap-style{
    label {
        font-size: 1.075em;
        font-weight: normal;
        color: black;
    }
    .name {
        font-weight: bold;
    }
}

.mapDescription {
    margin-top: 0.5em;
}
</style>
