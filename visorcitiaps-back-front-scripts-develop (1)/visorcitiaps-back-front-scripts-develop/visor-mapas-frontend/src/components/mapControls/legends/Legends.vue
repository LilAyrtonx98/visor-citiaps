<template>

<div style="background: white; z-index: 10;" v-show="isActive" :class="{widgetContainer : !isFixed, legendContainer : !isFixed}" :id="isFixed? 'legendIDFix' : 'legendID'">
  <div v-if="!isFixed" :class="{widgetHeader : !isFixed}" id="legendIDheader">
      <p style="font-weight: bold;">Simbología</p>
      
      <div>
        <button class="btn btn-danger" v-on:click="exitLegend()">
          <font-awesome-icon icon="times"></font-awesome-icon>
        </button>
      </div>

  </div>
  <div :class="{widgetBody : !isFixed}" v-show="!isHidden">
    <div :class="{resizable : !isFixed}" v-if="activeLayers.length > 0">
      
      <ListPlegableLegend
      v-for="(layer,i) in activeLayers" v-bind:key="i"
      v-bind:layer="layer"
      >
      </ListPlegableLegend>

    </div>
    <div :class="{resizable : !isFixed}" v-else>
      <span>No hay datos</span>
    </div>
  </div>
</div>

</template>

<script>
import ListPlegableLegend from '@/components/mapControls/legends/ListPlegableLegend.vue';
import {draggableDiv} from '@/components/mixin/draggableDiv.js'
import {vuex} from '@/components/mixin/vuex.js'

function arrayRemove(arr, value) {

   return arr.filter(function(ele){
       return ele != value;
   });

}

export default {
  name: 'HelloWorld',
  props: ['isActive', 'isFixed'],
  mixins:[draggableDiv, vuex],
  data() {
    return {
      isHidden: false,
    }
  },
  components:{
    ListPlegableLegend
  },
  computed: {

  },
  methods:{
    exitLegend: function() {
      this.$emit('onLegendAction','legend','exit', null)
    },
  },
  updated () {

  },
  created () {

  },
  mounted () {
    this.makeDraggable("legendID");
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

.legendContainer{
  left: 78vw;
  top: 4em;
}

.resizable{
  margin-top: 0.5em;
  height: 20vh;
  width: 5vw;
  min-height: 20vh;
  min-width: 20vw;
  resize: vertical;
  overflow: auto;
}

.scrollable{
  height: 110px;
  overflow-y: scroll;
}

.widget-container-alt {
	position: relative;
}
  
</style>
