<template>
<main>
    <h2>Visor de mapas</h2>

    <div class="container-1-1-1">

        <Card
        v-for="map in mapsAssignedToUserID"
        v-bind:key="map.index"
        v-bind:imgsrc="map.imgurl? map.imgurl : '/images/forest-background.jpg'"
        v-bind:imgalt="map.name"
        v-bind:title="map.name"
        v-bind:subtitle="map.desc"
        v-bind:url="'/visor/' + map.id"
        >
        </Card>

    </div>
</main>
</template>

<style>

</style>

<script>
import LogItem from '@/components/LogItem.vue'
import Card from '@/components/Card.vue'

import {restVisor} from '@/components/mixin/restVisor.js'
import {vuex} from '@/components/mixin/vuex.js'

export default{
  name: 'home',
  mixins: [restVisor, vuex],
  data(){
    return{
        logs : [],
        notifications: []
    }
  },
  components:{
      LogItem,
      Card
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
    this.getMapsAssigned()
  }
  
}
</script>

<style lang="scss" scoped>

</style>
