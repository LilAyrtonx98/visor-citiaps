<template>
  <div class="statContainer" :class="{statContainerAlt : chartContainerAlt}">
    <div v-if="type==0">
      <p class="stat-title">No hay datos</p>
      <p class="stat-subtitle">Selecciona un área de la capa para visualizar los datos correspondientes</p>
    </div>
    <div v-else-if="type==2">
      <p class="stat-title">No hay resultados</p>
      <p class="stat-subtitle">Ejecuta un geoproceso para visualizar los datos resultantes</p>
    </div>
    <div v-else>
      <p class="stat-title">No se ha seleccionado una capa</p>
      <p class="stat-subtitle">Selecciona una capa para poder ver los datos de un área</p>
    </div>
  </div>
  
</template>

<script>
export default{
  props:['type', 'chartContainerAlt'],
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

<style lang="scss" scoped>

</style>