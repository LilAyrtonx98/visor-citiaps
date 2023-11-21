<template>
<div class="screen-background" v-if="show">

    <div class="lds-dual-ring"></div>

</div>
</template>

<style lang="scss" scoped>

.screen-background {
  display: flex;
  position: fixed; /* Stay in place */
  z-index: 1; /* Sit on top */
  //padding-top: 100px; /* Location of the box */
  left: 0;
  top: 0;
  width: 100%; /* Full width */
  height: 100%; /* Full height */
  overflow: auto; /* Enable scroll if needed */
  background-color: rgb(0,0,0); /* Fallback color */
  background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
}

/*
  Spinner
*/
.lds-dual-ring {
  display: inline-block;
  width: 64px;
  height: 64px;
  margin: auto;
}
.lds-dual-ring:after {
  content: " ";
  display: block;
  width: 46px;
  height: 46px;
  margin: 1px;
  border-radius: 50%;
  border: 5px solid #03440c;
  border-color: #03440c transparent #03440c transparent;
  animation: lds-dual-ring 1.2s linear infinite;
}
@keyframes lds-dual-ring {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}


</style>

<script>
import {vuex} from '@/components/mixin/vuex.js'

export default{
  props:[],
  mixins: [vuex],
  data(){
    return {
      show: false,
    }
  },
  computed: {
    isActive: function() {
      return this.mapControlValues('loadScreen').active
    }
  },
  methods: {

  },
  created(){
  
  },
  watch: {
  'isActive': function (newValue, oldValue) {
      this.show = newValue
    },
  },

}
</script>