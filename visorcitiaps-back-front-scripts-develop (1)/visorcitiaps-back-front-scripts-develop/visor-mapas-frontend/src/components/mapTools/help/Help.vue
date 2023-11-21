<template>
  <modal name="Help" style="z-index: 100000001" :height="'auto'" :width="'90%'">
      <div class="modal-content">

          <div class="modal-item container-1">                

                <h2>Atajos de teclado</h2>
                
                <div class="button-explanation ">
                    
                    <div class="key">
                        ctrl
                    </div>
                    <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
                    <div class="key">
                        A
                    </div>
                    <p>Muestra los atajos del teclado en el mapa</p>
                </div>

                <div class="button-explanation ">
                    <div class="key">
                        <span style="font-size: 0.85em;">shift</span>
                    </div>
                    <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
                    <div class="key">
                        <font-awesome-icon icon="arrow-up"/>
                    </div>
                    <p>Aumenta el zoom al mapa</p>
                </div>

                <div class="button-explanation ">
                    <div class="key">
                        <span style="font-size: 0.85em;">shift</span>
                    </div>
                    <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
                    <div class="key">
                        <font-awesome-icon icon="arrow-down"/>
                    </div>
                    <p>Reduce el zoom al mapa</p>
                </div>

                <div class="button-explanation ">
                    <div class="key">
                        ctrl
                    </div>
                    <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
                    <div class="key">
                        S
                    </div>
                    <p>Muestra la simbología de las capas en el mapa</p>
                </div>

                <div class="button-explanation ">
                    <div class="key">
                        ctrl
                    </div>
                    <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
                    <div class="key">
                        U
                    </div>
                    <p>Muestra la ubicación del usuario en el mapa</p>
                </div>

                <div class="button-explanation ">
                    <div class="key">
                        ctrl
                    </div>
                    <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
                    <div class="key">
                        P
                    </div>
                    <p>Carga el mapa en toda la pantalla</p>
                </div>

                <div class="button-explanation ">
                    <div class="key">
                        ctrl
                    </div>
                    <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
                    <div class="key">
                        H
                    </div>
                    <p>Abre la caja de herramientas</p>
                </div>
              
          </div>

          <div class="buttons-container modal-buttons" >
              <button v-on:click="hideModal()" class="btn btn-secondary">Volver</button>
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

    }
  },
  computed: {
      isActive: function() {
          return this.mapToolValues('help').active
      }
  },
  methods: {
    showModal: function(){
        this.$modal.show("Help")
    },
    hideModal: function(){
        this.$modal.hide("Help")
    }
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
p {
    margin-left: 1em;
}
.button-explanation{
    display: flex;
    flex-direction: row;
    width: 100%;
    align-items: center;
}
.key{
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
