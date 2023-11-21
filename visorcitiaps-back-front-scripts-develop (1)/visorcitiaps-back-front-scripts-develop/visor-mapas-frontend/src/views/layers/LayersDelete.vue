<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>{{layer.name}}</h2>
    </div>

    <div class="container-body">

        <span>¿Realmente deseas eliminar esta capa?</span>
        <br>
        <span>Si lo haces, un registro de este recurso quedará en la base de datos</span>
        <br><br>
        <button class="btn btn-danger" style="margin-right: 1em;" v-on:click="deleteLayer()">Sí</button>
        <button class="btn btn-secondary" v-on:click="onClickDontDelete()">No</button>

    </div>

</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'

import {restLayer} from '@/components/mixin/restLayer.js'

export default{
    name: 'test',
    mixins: [restLayer],
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
              {name:'Capas', path: '/capas' },
              {name: this.layer.name, path: '/capas/' + this.$route.params.idLayer},
              {name: 'Eliminar capa'}
              ]
        }
    },
    props: [],
    methods:{
        onClickDontDelete: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/capas/' + this.$route.params.idLayer)
        },
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.layers? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getLayer();
    }
}
</script>

<style lang="scss" scoped>

</style>
