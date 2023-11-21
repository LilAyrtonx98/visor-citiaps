<template>
  <modal name="Export" style="z-index: 100000001" :height="'auto'" :width="'90%'">
      <div class="modal-content">
            <div class="modal-item container-1">                
                <h2>Exportar Mapa</h2>
                <label for="mi-select">Formato</label>
                <select v-model="exportFormat">
                    <option
                    value="image">
                        Imagen PNG
                    </option>
                    <option
                    value="document">
                        Documento PDF
                    </option>
                </select>
                <label for="name">Nombre</label>
                <input v-model="filename" type="text" id="name" placeholder="Ingresa un nombre"/>
                <p v-if="exportFormat=='image'">
                Guarda el mapa como una imange PNG. No incluye la simbología.
                </p>
                <p v-else>
                Guarda el mapa como un documento PDF. Incluye la simbología.
                </p>
            </div>

          <div class="buttons-container modal-buttons" >
              <button v-on:click="hideModal()" class="btn btn-secondary">Cancelar</button>
              <button v-on:click="onClickOK()" class="btn btn-primary">Exportar</button>
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
      exportFormat: '',
      filename: ''
    }
  },
  computed: {
      isActive: function() {
          return this.mapToolValues('export').active
      }
  },
  methods: {
    showModal: function(index){
        this.$modal.show("Export")
    },
    hideModal: function(){
        this.$modal.hide("Export")
    },
    onClickOK: function() {
        this.executeMapToolAction('export', 'exportMap', {format: this.exportFormat, filename: this.filename})
        this.$modal.hide("Export")
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
