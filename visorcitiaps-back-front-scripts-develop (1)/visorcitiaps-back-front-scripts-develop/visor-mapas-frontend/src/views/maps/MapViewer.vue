<template>
<MapContainer>

</MapContainer>
</template>

<style lang="scss" scoped>

</style>

<script>
import MapContainer from '@/components/MapContainer.vue'

// Mixin
import {restMap} from '@/components/mixin/restMap.js'
import {vuex} from '@/components/mixin/vuex.js'

export default{
  name: 'home',
  mixins: [restMap, vuex],
  data(){
    return{
      
      }
  },
  components:{
      MapContainer
  },
  computed: {
  },
  props: [],
  methods:{

  },
  updated () {
    
  },
  created () {
    //Verificar permisos
    this.$store.getters.getPermissions? true:this.$router.push('/inicio')
    this.$store.getters.getPermissions.maps? true:this.$router.push('/pagina-no-encontrada')
  },
  mounted () {
    // Obtener capas y geoprocesos asignados al mapa
    this.getLayersByMap();
    this.getGeoprocessingsByMap();
  },
  
}
</script>
