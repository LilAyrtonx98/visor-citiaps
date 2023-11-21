<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Nuevo grupo</h2>
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
            <button class="btn btn-secondary" v-on:click="onClickGoNext()">Crear grupo</button>
        </div>

    </div>

    <div class="container-list"> 
        
    </div>
 
</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'

export default{
    name: 'test',
    data(){
        return{
            group: {
                name: '',
                desc: '',
            },
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
                {name: 'Nuevo grupo'},
          ]
        }
    },
    props: [],
    methods:{
        onClickGoBack: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/usuarios/grupos-de-usuarios')
        },
        onClickGoNext: function() {
            var that = this;
            this.$validator.validateAll()
            .then(function (response) {
                if(response) {
                    console.log('OK')
                    that.postGroup();
                } else {
                    console.log('NOPE')
                }
            })
        },
        postGroup: function(){
            var that = this;
            this.$http.post(this.$store.getters.backendurl + 'api/v1/groups', 
            this.group
            )
            .then(function (response) {
                //Obtener ID de la tarea recién creada y pasar a la vista de asignar usuarios
                let idGeo = response.data.id
                that.$router.push('/usuarios/grupos-de-usuarios/'+ idGeo)
            })
            .catch(function (error) {
                console.log(error);
            });
        }
    },
    updated () {

    },
    created () {
        //Verificar permisos
        this.$store.getters.getPermissions.users? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {

    }
}
</script>

<style lang="scss" scoped>

</style>
