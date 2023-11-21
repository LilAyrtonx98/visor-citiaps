<template>
<main class="container-1">

    <Login
    >
    </Login>

</main>
</template>

<style>

</style>

<script>
import Login from '@/components/Login.vue'

export default{
  name: 'home',
  data(){
    return{
      
    }
  },
  components:{
    Login
  },
  computed: {

  },
  props: [],
  methods:{
    hello: function() {
      console.log("Hello");
    },
  },
  updated () {

  },
  created () {
        //Verificar permisos
        this.$store.getters.getPermissions? this.$router.push('/dashboard'):false
  },
  mounted () {

  }
  
}
</script>

<style lang="scss" scoped>
body main {
	margin: 0 0;
    padding: 0 0;
    align-content: center;
    justify-items: center;

    background-image: url('/images/forest-background.jpg');
    background-size: cover;
    background-position-y: center;
    background-position-x: center;
    //filter: grayscale(10%);
    filter: saturate(1.25);
}
</style>