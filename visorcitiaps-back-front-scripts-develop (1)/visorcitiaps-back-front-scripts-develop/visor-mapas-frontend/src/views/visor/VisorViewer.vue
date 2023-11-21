<template>
<MapContainer>

</MapContainer>
</template>

<style lang="scss" scoped>

</style>

<script>
import MapContainer from '@/components/MapContainer.vue'

// Mixin
import {restVisor} from '@/components/mixin/restVisor.js'
import {vuex} from '@/components/mixin/vuex.js'

export default{
  name: 'home',
  mixins: [restVisor, vuex],
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
    this.$store.getters.getPermissions.visor? true:this.$router.push('/pagina-no-encontrada')
  },
  mounted () {
    // Obtener capas y geoprocesos disponibles para el usuario
    this.getLayersAssigned(this.$route.params.idMap);
    this.getGeoprocessingsAssigned(this.$route.params.idMap);
    
  },
  
}
</script>
