<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>{{user.firstname + ' ' + user.lastname}}</h2>
    </div>

    <div class="container-body">

        <span>¿Realmente deseas eliminar este usuario?</span>
        <br>
        <span>Si lo haces, un registro de este recurso quedará en la base de datos</span>
        <br><br>
        <button class="btn btn-danger" style="margin-right: 1em;" v-on:click="deleteUser()">Sí</button>
        <button class="btn btn-secondary" v-on:click="onClickDontDelete()">No</button>

    </div>

</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'

import {restUser} from '@/components/mixin/restUser.js'

export default{
    name: 'test',
    mixins: [restUser],
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
              {name:'Usuarios', path: '/usuarios' },
              {name: this.user.firstname + ' ' + this.user.lastname, path: '/usuarios/' + this.$route.params.idUser},
              {name: 'Editar usuario'}
              ]
        }
    },
    props: [],
    methods:{
        onClickDontDelete: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/usuarios/' + this.$route.params.idUser)
        },
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.users? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getUser();
    }
}
</script>

<style lang="scss" scoped>

</style>
