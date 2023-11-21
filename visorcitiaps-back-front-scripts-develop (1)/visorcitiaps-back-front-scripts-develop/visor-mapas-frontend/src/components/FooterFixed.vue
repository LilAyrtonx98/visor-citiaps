<template>
  <footer>
    <div class="watermark">
      <router-link v-if="$store.getters.getPermissions" style="color: white;" :to="'/ayuda'">{{'¿Necesitas ayuda?'}}</router-link>
    </div>
    <div></div>
    <img alt="Logo CITIAPS" src="../assets/img/imago-citiaps-blanco.svg"/>
    <img alt="Logo USACH" src="../assets/img/imago-usach-blanco.svg"/>
  </footer>
</template>

<script>
export default {
  name: 'FooterFixes',
  props: {
    msg: String
  },
  methods: {
    toggleFooter() {
      
    }
  },
  mounted() {
    
  },
  beforeDestroyed() {

  },
  watch: {
    
  }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
