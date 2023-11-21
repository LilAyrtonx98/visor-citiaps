<template>
<div v-if="logs" class="list-item-area">
    <template 
    v-for="(item) in logs" >

      <p class="list-item-sub-1" style="font-size: 1.1em;" v-bind:key="item.index">{{item.action}}</p>
      <p class="list-item-sub-1" v-bind:key="item.index" v-if="fullMode">{{item.id_resource}}</p>        
      <label v-bind:key="item.index" class="list-item-sub-1" style="grid-column-end: 3; margin-top: 0em; align-self: end;">{{dateFormatter(item.created_at)}}</label>
      
      <span v-bind:key="item.index" class="list-item-separator"></span>
      
    </template>
    <span v-if="logs.length<=0">No hay registros</span>
</div>
</template>

<style lang="scss" scoped>
.btnSelected{
  background-color: #f7e200;
}
.btnSelected:hover{
  background-color: #f7e200;
}
</style>

<script>
var moment = require('moment');

export default{
  props:['logs', 'fullMode'],
  data(){
    return {
        selectedID: null
    }
  },
  computed: {

  },
  methods: {
    dateFormatter: function (date) {
          moment.locale('es');
          return moment(date).format('LL');
      },
    onClickSelectItem: function (item) {
      this.selectedID = item.id;
      this.$emit('itemSelected', item);
    },
    isSelected: function (id) {
        if(this.selectedID == id) {
          return true;
        }
        return false;
    }
  },
  created(){
  
  },
  mounted(){

  },
  watch: {

  }
}
</script>