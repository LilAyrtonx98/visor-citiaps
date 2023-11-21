<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>{{group.name}}</h2>
    </div>

    <div class="container-body">

        <span>¿Realmente deseas eliminar este grupo?</span>
        <br>
        <span>Si lo haces, un registro de este recurso quedará en la base de datos</span>
        <br><br>
        <button class="btn btn-danger" style="margin-right: 1em;" v-on:click="deleteGroup()">Sí</button>
        <button class="btn btn-secondary" v-on:click="onClickDontDelete()">No</button>

    </div>

</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'

import {restGroup} from '@/components/mixin/restGroup.js'

export default{
    name: 'test',
    mixins: [restGroup],
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
                {name: 'Usuarios', path: '/usuarios'},
                {name: 'Grupos', path: '/usuarios/grupos-de-usuarios/'},
                {name: this.group.name, path: '/usuarios/grupos-de-usuarios/' + this.$route.params.idGroup},
                {name: 'Eliminar grupo'}
          ]
        }
    },
    props: [],
    methods:{
        onClickDontDelete: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/usuarios/grupos-de-usuarios/' + this.$route.params.idGroup)
        },        
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.users? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getGroup();
    }
}
</script>

<style lang="scss" scoped>

</style>
