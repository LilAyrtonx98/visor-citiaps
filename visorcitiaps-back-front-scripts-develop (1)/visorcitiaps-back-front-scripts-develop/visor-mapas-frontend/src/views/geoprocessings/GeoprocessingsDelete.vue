<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>{{geo.name}}</h2>
    </div>

    <div class="container-body">

        <span>¿Realmente deseas eliminar este geoproceso?</span>
        <br>
        <span>Si lo haces, un registro de este recurso quedará en la base de datos</span>
        <br><br>
        <button class="btn btn-danger" style="margin-right: 1em;" v-on:click="deleteGeoprocessing()">Sí</button>
        <button class="btn btn-secondary" v-on:click="onClickDontDelete()">No</button>

    </div>

</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'

import {restGeoprocessing} from '@/components/mixin/restGeoprocessing.js'

export default{
    mixins: [restGeoprocessing],
    data(){
        return{

        }
    },
    components:{
        Breadcrumb,
    },
    computed: {
        routesBread: function(){
          return [
              {name:'Geoprocesos', path: '/geoprocesos' },
              {name: this.geo.name, path: '/geoprocesos/' + this.$route.params.idGeo},
              {name: 'Eliminar geoproceso'}
              ]
        }
    },
    props: [],
    methods:{
        onClickDontDelete: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/geoprocesos/' + this.$route.params.idGeo)
        },
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.geo? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getGeoprocessing();
    }
}
</script>

<style lang="scss" scoped>

</style>
