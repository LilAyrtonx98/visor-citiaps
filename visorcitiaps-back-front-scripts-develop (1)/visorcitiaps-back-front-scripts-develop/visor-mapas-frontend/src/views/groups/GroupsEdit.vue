<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Editar grupo</h2>
    </div>

    <div class="container-action">
        
    </div>

    <div class="container-body">
        <label for="name">Nombre</label>
        <input id="name" v-model="group.name" type="text"
        name="nombre" :class="{'error' : errors.has('nombre') }" v-validate="'required'"
        :placeholder="[[errors.has('nombre') ? errors.first('nombre') : 'Ingresa el nombre de la categoría']]"/>

        <label for="desc">Descripción</label>
        <textarea id="desc" v-model="group.desc" placeholder="Ingresa la descripción de la categoría"></textarea>


        <div class="container-action-left" style="justify-content: start;">
            <button class="btn btn-secondary" v-on:click="onClickGoBack()">Volver</button>
            <button class="btn btn-secondary" v-on:click="onClickGoNext()">Actualizar categoría</button>
        </div>

    </div>

    <div class="container-list"> 
        
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
        Breadcrumb
    },
    computed: {
        routesBread: function(){
          return [
                {name: 'Usuarios', path: '/usuarios'},
                {name: 'Grupos', path: '/usuarios/grupos-de-usuarios/'},
                {name: this.group.name, path: '/usuarios/grupos-de-usuarios/' + this.$route.params.idGroup},
                {name: 'Editar grupo'}
          ]
        }
    },
    props: [],
    methods:{
        onClickGoBack: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/usuarios/grupos-de-usuarios/' + this.$route.params.idGroup)
        },
        onClickGoNext: function() {
            var that = this;
            this.$validator.validateAll()
            .then(function (response) {
                if(response) {
                    console.log('OK')
                    that.putGroup();
                } else {
                    console.log('NOPE')
                }
            })
        }
    },
    updated () {

    },
    created () {
        //Verificar permisos
        this.$store.getters.getPermissions.users? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getGroup();
    }
}
</script>

<style lang="scss" scoped>

</style>
