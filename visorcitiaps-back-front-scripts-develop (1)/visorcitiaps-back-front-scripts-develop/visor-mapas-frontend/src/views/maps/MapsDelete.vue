<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>{{map.name}}</h2>
    </div>

    <div class="container-body">

        <span>¿Realmente deseas eliminar este mapa?</span>
        <br>
        <span>Si lo haces, un registro de este recurso quedará en la base de datos</span>
        <br><br>
        <button class="btn btn-danger" style="margin-right: 1em;" v-on:click="deleteMap()">Sí</button>
        <button class="btn btn-secondary" v-on:click="onClickDontDelete()">No</button>

    </div>

</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'

import {restMap} from '@/components/mixin/restMap.js'

export default{
    name: 'test',
    mixins: [restMap],
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
                {name: 'Mapas', path: '/mapas'},
                {name: this.map.name, path: '/mapas/' + this.$route.params.idMap},
                {name: 'Eliminar mapa'}
          ]
        }
    },
    props: [],
    methods:{
        onClickDontDelete: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/mapas/' + this.$route.params.idMap)
        }
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
