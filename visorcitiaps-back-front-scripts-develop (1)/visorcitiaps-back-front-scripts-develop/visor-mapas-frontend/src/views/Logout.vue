<template>

</template>

<script>

export default{
  props:['logs'],
  data(){
    return {
        
    }
  },
  computed: {

  },
  methods: {
    onClickLogout: function(){
      /*
      //Quitar permisos de usuario
      document.cookie = "session_public" + '=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';
      this.$store.commit('setPermissions');
      //Quitar token
      document.cookie = "session_token" + '=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';
      this.$store.commit('setToken');
      //Inicio
      this.$router.push('/inicio')
      */
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/logout')
        .then(function (response) {
            that.message = response.data.message;
            //Quitar permisos de usuario
            document.cookie = "session_public" + '=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';
            that.$store.commit('setPermissions');
            that.$router.push('/inicio')
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
    }
  },
  created(){
    this.onClickLogout();
  },
  watch: {

  }
}
</script>