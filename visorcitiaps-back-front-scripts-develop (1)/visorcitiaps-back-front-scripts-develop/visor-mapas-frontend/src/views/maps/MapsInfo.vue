<template>
<main class="container-right-nav">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>{{map.name}}</h2>
    </div>

    <div class="container-body">
        <strong>Información</strong>
        <p>{{dateFormatter(map.created_at)}}</p>
        <p>{{map.desc}}</p>
        <br>

    </div>


    <SideNav
    class="container-sidenav"
    v-bind:links="routesNav"
    >
    </SideNav>

</main>
</template>

<script>
// @ is an alias to /src
import SideNav from '@/components/SideNav.vue'
import Breadcrumb from '@/components/Breadcrumb.vue'
import Pagination from '@/components/Pagination.vue'

import {formatter} from '@/components/mixin/formatter.js'
import {restMap} from '@/components/mixin/restMap.js'

export default{
    name: 'test',
    mixins: [formatter, restMap],
    data(){
        return{

        }
    },
    components:{
        SideNav,
        Breadcrumb
    },
    computed: {
        routesNav: function(){
          return [
                {name:'Mapas'},
                {name:'Ver mapa' , path: '/mapas/' + this.$route.params.idMap + '/visor'},
                {name:'Usuarios asignados' , path: '/mapas/' + this.$route.params.idMap + '/asignar-usuarios'},
                {name:'Capas asignadas' , path: '/mapas/' + this.$route.params.idMap + '/asignar-capas'},
                {name:'Geoprocesos asignados' , path: '/mapas/' + this.$route.params.idMap + '/asignar-geoprocesos'},
                {name:'Editar mapa' , path: '/mapas/' + this.$route.params.idMap + '/editar-mapa'},
                {name:'Historial de cambios' , path: '/mapas/' + this.$route.params.idMap + '/historial-de-cambios'},
                {name:'Eliminar mapa' , path: '/mapas/' + this.$route.params.idMap + '/eliminar-mapa'},
              ]
        },
        routesBread: function(){
          return [
                {name: 'Mapas', path: '/mapas'},
                {name: this.map.name},
          ]
        }
    },
    props: [],
    methods:{

    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.maps? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getMap();
    }
}
</script>

<style lang="scss" scoped>

</style>
