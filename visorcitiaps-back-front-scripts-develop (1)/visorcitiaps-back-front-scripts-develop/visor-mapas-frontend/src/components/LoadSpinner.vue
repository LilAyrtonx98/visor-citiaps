<template>
<div class="lds-dual-ring"></div>
</template>

<style lang="scss" scoped>

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
export default{
  props:['show'],
  data(){
    return {

    }
  },
  computed: {

  },
  methods: {

  },
  created(){
  
  },
  watch: {
    
  }
}
</script>