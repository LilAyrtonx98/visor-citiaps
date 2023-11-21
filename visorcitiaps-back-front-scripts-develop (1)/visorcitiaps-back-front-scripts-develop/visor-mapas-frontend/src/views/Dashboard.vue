<template>
<main>
    <h2>Resumen general</h2>

    <div class="container-1-1-1">

        <div
        v-if="$store.getters.getPermissions.maps"
        class="counter"
        >
            <span style="font-size: 1.5em;">{{mapsCount}}</span>
            <span>Mapas</span>
        </div>

         <div
        v-if="$store.getters.getPermissions.layers"
        class="counter"
        >
            <span style="font-size: 1.5em;">{{categoriesCount}}</span>
            <span>Categorías de capas</span>
        </div>

        <div
        v-if="$store.getters.getPermissions.users"
        class="counter"
        >
            <span style="font-size: 1.5em;">{{groupsCount}}</span>
            <span>Grupos de usuarios</span>
        </div>

        <div
        v-if="$store.getters.getPermissions.geo"
        class="counter"
        >
            <span style="font-size: 1.5em;">{{geoCount}}</span>
            <span>Geoprocesos</span>
        </div>
        
        <div
        v-if="$store.getters.getPermissions.layers"
        class="counter"
        >
            <span style="font-size: 1.5em;">{{layersCount}}</span>
            <span>Capas</span>
        </div>


        <div
        v-if="$store.getters.getPermissions.users"
        class="counter"
        >
            <span style="font-size: 1.5em;">{{usersCount}}</span>
            <span>Usuarios</span>
        </div>

    </div>

    <br>

    <h2>Resumen usuario</h2>

    <div class="container-1">

        <div
        class="counter"
        >
            <span style="font-size: 1.5em;">{{mapsAssignedToUserID.length}}</span>
            <span>Mapas asignados al usuario</span>
        </div>

    </div>


</main>
</template>

<style>

</style>

<script>
import LogItem from '@/components/LogItem.vue'

import {vuex} from '@/components/mixin/vuex.js'
import {restCount} from '@/components/mixin/restCount.js'
import {restVisor} from '@/components/mixin/restVisor.js'

export default{
  name: 'home',
  mixins: [vuex, restCount, restVisor],
  data(){
    return{
        
    }
  },
  components:{
      LogItem
  },
  computed: {

  },
  props: [],
  methods:{

  },
  updated () {
      console.log(this.$store.getters.getMapSelectedTool)
  },
  created () {
    //Verificar permisos
    this.$store.getters.getPermissions? true:this.$router.push('/inicio')
    this.$store.getters.getPermissions.visor? true:this.$router.push('/pagina-no-encontrada')
  },
  mounted () {
      this.getMapsCount()
      this.getLayersCount()
      this.getGeoCount()
      this.getUsersCount()
      this.getCategoriesCount()
      this.getGroupsCount()

      this.getMapsAssigned()
  }
  
}
</script>

<style lang="scss" scoped>
.counter{
    height: 10vh;
    border-radius: 6px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}
.counter:nth-child(odd) {
    background-color: rgb(3,68,12);
    color: white;
}
.counter:nth-child(even) {
    background-color: rgba(51, 51, 51, 0.5);
    color: black;
}
</style>
