<template>
<div style="width: 100%">
    <modal name="addAnnotation" style="z-index: 100000001" :height="'auto'" :width="'100%'">
      <div class="modal-content">
        <div class="modal-item container-1">
          <h2>Agregar anotación</h2>
          <label for="annotation-text">Texto</label>
          <textarea v-model="annotationData.text" id="annotation-text"></textarea>
          <div>
          <input v-model="annotationData.is_shared" id="share" type="checkbox"> 
          <label for="share">Compartir con grupo</label>
          </div>
        </div>
          <div class="buttons-container modal-buttons" >
              <button v-on:click="hideModal()" class="btn btn-secondary">Cancelar</button>
              <button v-on:click="onClickOk()" class="btn btn-primary">Confirmar</button>
          </div>
      </div>
    </modal>
    <AnnotationInfo :info="annotationInfo" @saveEdit="editAnnotationsLayers()"></AnnotationInfo>
    <div v-show="isActive" class="widgetContainer" id="annotationsID">
      <div class="widgetHeader">
        <p style="font-weight: bold">Anotaciones</p>
        <div>
          <button class="btn btn-info" v-on:click="isHidden = !isHidden">
            <font-awesome-icon v-if="!isHidden" icon="minus"></font-awesome-icon>
            <font-awesome-icon v-else icon="plus"></font-awesome-icon>
          </button>
          <button class="btn btn-danger" v-on:click="exitAnnotations()">
            <font-awesome-icon icon="times"></font-awesome-icon>
          </button>
        </div>
      </div>
      <div class="widgetBody" v-show ="!isHidden">
        <input id="showAnnotations1" type="checkbox" v-model="showUserAnnotations">
        <label for="showAnnotations1">Mis anotaciones <span class="dotUserAnnotation"></span></label>
        <br>
        <input id="showAnnotations2" type="checkbox" v-model="showGroupAnnotations">
        <label for="showAnnotations2">Anotaciones del grupo <span class="dotGroupAnnotation"></span></label>
        <br>
        <br>
        <button v-on:click="addAnnotation()" class="btn btn-secondary">
          <span>Agregar anotación</span>
        </button>
      </div>
    </div>
</div>
</template>

<script>
import {draggableDiv} from '@/components/mixin/draggableDiv.js'
import {vuex} from '@/components/mixin/vuex.js'
import AnnotationInfo from '@/components/mapTools/annotations/AnnotationInfo.vue'
export default {
  name: 'annotations',
  mixins:[vuex, draggableDiv],
  props: ['isFixed', 'currentMode'],
  components:{
    AnnotationInfo
  },
  data(){
    return{
      isHidden: false,
      annotationData:{
        "id_user": "",
        "id_map": "",
        "id_group": "",
        "text":"",
        "location":{
          "type": "Point",
          "coordinates":[0,0]
        }, 
        "is_shared":false
      },
      myself: null,
      showUserAnnotations: true,
      showGroupAnnotations: true,
      annotationInfo: ""
    }
  },
  methods:{
    exitAnnotations: function(){
      this.executeMapToolAction('annotations','removeListener')
      this.executeMapToolAction('annotations','removeAnnotationsByUser')
      this.executeMapToolAction('annotations','removeAnnotationsByGroup')
      // Cerrar la herramienta
      this.selectMapTool('annotations', false)
    },
    addAnnotation: function(){
      this.executeMapToolAction('annotations','addListener')
    },
    showModal: function(){
      this.cleanAnnotationData()
      this.$modal.show('addAnnotation')
    },
    hideModal: function(){
      this.executeMapToolAction('annotations','removeListener')
      this.$modal.hide('addAnnotation')
    },
    onClickOk: function(){
      this.insertAnnotation()
      this.executeMapToolAction('annotations','removeListener')
      this.$modal.hide("addAnnotation")
    },
    getMyself: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/accounts/me')
        .then(function (response) {
            that.myself = response.data;
        })
        .catch(function (error) {
            console.log(error)
        });
    },
    insertAnnotation: function(){
      var that = this
      console.log("anotacion", this.annotationData)
      this.$http.post(this.$store.getters.backendurl + 'api/v1/annotations', this.annotationData)
      .then(function(response){
        that.executeMapToolAction('annotations', 'removeAnnotationsByUser')
        that.getAnnotationsByUser(that.myself.id)
        console.log(response)
      })
      .catch(function (error) {
          console.log(error);
      });
    },
    cleanAnnotationData(){
      this.annotationData.text=""
      this.annotationData.is_shared=false
    },
    getAnnotationsByUser: function(userId){
        var that=this
        this.$http.get(this.$store.getters.backendurl + 'api/v1/annotations/'+userId)
        .then(function (response){
          if(response.data) that.executeMapToolAction('annotations','showAnnotationsByUser', {annotationsByUser: response.data})
        })
        .catch(function(error){
            console.log(error)
        })
    },
    getAnnotationsByGroup: function(groupId){
        var that=this
        this.$http.get(this.$store.getters.backendurl + 'api/v1/annotations/'+groupId+'/group')
        .then(function (response){
          if(response.data) that.executeMapToolAction('annotations','showAnnotationsByGroup', {annotationsByGroup: response.data, userId: that.myself.id})
        })
        .catch(function(error){
            console.log(error)
        })
    },
    editAnnotationsLayers: function(){
        this.executeMapToolAction('annotations', 'removeAnnotationsByUser')
        this.getAnnotationsByUser(this.myself.id)
        this.executeMapToolAction('annotations', 'removeAnnotationsByGroup')
        this.getAnnotationsByGroup(this.myself.id_group)
    }
  },
  computed: {
    isActive: function() {
      return this.isFixed? this.isMobile : this.mapToolValues('annotations').active && !this.isMobile
    }
  },
  watch:{
    'mapToolExecuted': function(){
      if(this.mapToolValues('annotations').action == 'showModal'){
        this.annotationData.id_user=this.myself.id
        this.annotationData.id_group=this.myself.id_group
        this.annotationData.id_map=this.$route.params.idMap
        this.annotationData.location.coordinates=this.mapToolValues('annotations').data.annotationCoor
        this.showModal()
      }
      if(this.mapToolValues('annotations').action == 'showAnnotationInfo'){
        this.annotationInfo=this.mapToolValues('annotations').data.info
      }
    },
    'isActive': function(value){
      if(value){
        this.getAnnotationsByUser(this.myself.id)
        this.getAnnotationsByGroup(this.myself.id_group)
        this.executeMapToolAction('annotations', 'addAnnotationSelector')
      }
    },
    'showUserAnnotations': function(value){
      if(value){
        this.getAnnotationsByUser(this.myself.id)
      }
      else{
        this.executeMapToolAction('annotations','removeAnnotationsByUser')
      }
    },
    'showGroupAnnotations': function(value){
      if(value){
        this.getAnnotationsByGroup(this.myself.id_group)
      }
      else{
        this.executeMapToolAction('annotations', 'removeAnnotationsByGroup')
      }
    }
  },
  created: function(){
    this.getMyself()
  },
  mounted: function(){
    this.makeDraggable("annotationsID");
  }
}
</script>