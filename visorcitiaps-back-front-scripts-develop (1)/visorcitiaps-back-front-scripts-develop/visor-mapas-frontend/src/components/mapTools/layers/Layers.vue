<template>
  <div v-show="isActive" :class="{widgetContainer : !isFixed}" :id="isFixed? 'layersIDFix' : 'layersID'">
    <div v-if="!isFixed" :class="{widgetHeader : !isFixed}" id="layersIDheader">
        <p style="font-weight: bold;">Lista de capas</p>
        <div>
          <button  class="btn btn-primary" v-on:click="onClickModalNewLayer()">Añadir capa</button>

          <button class="btn btn-info" v-on:click="isHidden = !isHidden">
            <font-awesome-icon v-if="!isHidden" icon="minus"></font-awesome-icon>
            <font-awesome-icon v-else icon="plus"></font-awesome-icon>
          </button>
          <button class="btn btn-danger" v-on:click="exitLayers()">
            <font-awesome-icon icon="times"></font-awesome-icon>
          </button>
        </div>
        
    </div>
    <div :class="{widgetBody : !isFixed}" v-show="!isHidden">
      <div :class="{resizable : !isFixed}">
        <ListPlegable
        v-for="(category,i) in layers" v-bind:key="i"
        v-bind:category="category"
        v-bind:index="i"
        v-on:selectLayer="onClickSelectLayer"
        v-on:removeLayer="onClickRemoveLayer"
        v-on:layerAttributes="getLayerInfo"
        ></ListPlegable>
      </div>
    </div>
  </div>
</template>

<script>
import ListPlegable from '@/components/mapTools/layers/ListPlegable.vue'
import {draggableDiv} from '@/components/mixin/draggableDiv.js'
import {vuex} from '@/components/mixin/vuex.js'
import axios from 'axios'

export default {
  name: 'HelloWorld',
  props: ['isFixed'],
  mixins:[draggableDiv, vuex],
  data() {
    return {
      isHidden: false
    }
  },
  components:{
    ListPlegable,
  },
  computed: {
    isActive: function() {
      return this.isFixed? this.isMobile : this.mapToolValues('layers').active && !this.isMobile
    }
  },
  methods:{
    onClickSelectLayer: function(categoryIndex, layerIndex) {
      this.$store.dispatch('addActiveLayer', this.$store.getters.getLayers[categoryIndex].layers[layerIndex])
      this.executeMapToolAction('layers', 'addLayer', this.$store.getters.getLayers[categoryIndex].layers[layerIndex])
    },
    onClickRemoveLayer: function(categoryIndex, layerIndex) {
      this.$store.dispatch('removeActiveLayer', this.$store.getters.getLayers[categoryIndex].layers[layerIndex])
      console.log("removeLayer",this.$store.getters.getLayers[categoryIndex].layers[layerIndex])
      this.executeMapToolAction('layers', 'removeLayer', this.$store.getters.getLayers[categoryIndex].layers[layerIndex])
    },
    exitLayers: function() {
      this.selectMapTool('layers', false)
    },
    getLayerInfo: function(categoryIndex, layerIndex){
      var layer=this.$store.getters.getLayers[categoryIndex].layers[layerIndex]
      if(layer.provider.name=="arcgis"){
        this.arcgisFeatures(layer)
      }
      else this.geoserverFeatures(layer)
    },
    geoserverFeatures: function(layer){      
      var totalRecords=0
      var that=this
      var url=layer.provider.url+'&count=1&outputFormat=application%2Fjson'
      
      var reqgeo = new XMLHttpRequest();
      reqgeo.open('GET', url, false);
      reqgeo.withCredentials = false;
      reqgeo.send();

      if (reqgeo.status == 200){
        console.log("Success search geo query");
        var resgeo = JSON.parse(reqgeo.response);
        totalRecords=resgeo.totalFeatures
        that.executeMapToolAction('attributesTable', 'paginationInfo', {offset:0, totalRecords: totalRecords, layer:layer})
        that.selectMapTool('attributesTable', true)
      }else{
        console.log("Fail search geo query", reqgeo);
      }

      // DEPRECATED CORS ERROR needed withCredentials = false
      // this.$http.get(url)
      // .then(function(response){
      //     totalRecords=response.data.totalFeatures
      //     that.executeMapToolAction('attributesTable', 'paginationInfo', {offset:0, totalRecords: totalRecords, layer:layer})
      //     that.selectMapTool('attributesTable', true)
      // }, function(response){
      //   console.log("error", response)
      // })
    },
    arcgisFeatures: function(layer){      
      var totalRecords=0
      var subLayers=[]
      var that=this

      var req = new XMLHttpRequest();
      req.open('GET', layer.provider.url+'?f=json&pretty=true', false);
      req.withCredentials = false;
      req.send();

      if (req.status == 200){
        console.log("Success Attribute query 1");
        var resp = JSON.parse(req.response)
        console.log("sublayer", resp.layers);
        subLayers= resp.layers

        var req2 = new XMLHttpRequest();
        req2.open('GET', layer.provider.url+'/0/query?where=objectid>0&returnCountOnly=true&f=pjson', false);
        req2.withCredentials = false;
        req2.send();

        if(req2.status == 200){
          console.log("Success Attribute query 2");
          var resp2= JSON.parse(req2.response)
          totalRecords= resp2.count
          that.executeMapToolAction('attributesTable', 'paginationInfo', {offset:0, totalRecords: totalRecords, layer:layer, layerId: 0, subLayers: subLayers})
          that.selectMapTool('attributesTable', true)
          this.searchError = true;
          return;
        }else{
          console.log("Fail Attribute query 2");
        }
      }else{
        console.log("Fail Attribute query 1");
      }
      
      // DEPRECATED CORS ERROR needed withCredentials = false
      // that.$http.get(layer.provider.url+'?f=json&pretty=true')
      // .then(function(response){
      //   subLayers=response.data.layers
      //   that.$http.get.get(layer.provider.url+'/0/query?where=objectid>0&returnCountOnly=true&f=pjson')
      //   .then(function(response){
      //     totalRecords=response.data.count
      //     that.executeMapToolAction('attributesTable', 'paginationInfo', {offset:0, totalRecords: totalRecords, layer:layer, layerId: 0, subLayers: subLayers})
      //     that.selectMapTool('attributesTable', true)
      //   })
      // })
    },
    onClickModalNewLayer: function(){
      this.$router.push('/capas/nueva-capa')
    }
  },
  updated () {

  },
  created () {

  },
  mounted () {
    if(!this.isFixed) {
      this.makeDraggable("layersID");
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

.resizable{
  margin-top: 0.5em;
  height: 40vh;
  min-height: 20vh;
  min-width: 20vw;
  resize: both;
  overflow: auto;
}

.fixedWidget{
  margin-top: 0.5em;
  height: 100vh;
  min-height: 20vh;
  min-width: 20vw;
  resize: none;
  overflow: hidden
}

.scrollable{
  height: 110px;
  overflow-y: scroll;
}
</style>
