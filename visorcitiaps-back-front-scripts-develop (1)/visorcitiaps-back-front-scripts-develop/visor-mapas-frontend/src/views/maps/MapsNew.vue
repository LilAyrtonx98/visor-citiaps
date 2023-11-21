<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Nuevo mapa</h2>
    </div>

    <div class="container-action">
        
    </div>

    <div class="container-body">
        <label for="name">Nombre</label>
        <input id="name" v-model="map.name" type="text"
        name="nombre" :class="{'error' : errors.has('nombre') }" v-validate="'required'"
        :placeholder="[[errors.has('nombre') ? errors.first('nombre') : 'Ingresa el nombre del mapa']]"/>

        <label for="desc">Descripción</label>
        <textarea id="desc" v-model="map.desc" placeholder="Ingresa la descripción del mapa"></textarea>

        <label for="imgurl">Imagen</label>
        <input id="imgurl" v-model="map.imgurl" type="text" name="url imagen" placeholder='Ingresa la URL de la imagen'/>

        <div class="container-action-left" style="justify-content: start;">
            <button class="btn btn-secondary" v-on:click="onClickGoBack()">Volver</button>
            <button class="btn btn-secondary" v-on:click="onClickGoNext()">Crear mapa</button>
        </div>

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
            defaultImageURL: '/images/forest-background.jpg'
        }
    },
    components:{
        Breadcrumb
    },
    computed: {
        routesBread: function(){
          return [
                {name: 'Mapas', path: '/mapas'},
                {name:'Nuevo mapa'},
          ]
        }
    },
    props: [],
    methods:{
        onClickGoBack: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/mapas')
        },
        onClickGoNext: function() {
            var that = this;
            this.$validator.validateAll()
            .then(function (response) {
                if(response) {
                    if(that.map.imgurl !=null && !that.map.imgurl) {
                        that.map.imgurl = that.defaultImageURL
                    }
                    that.postMap();
                } else {

                }
            })
        }
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.maps? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {

    }
}
</script>

<style lang="scss" scoped>

</style>
