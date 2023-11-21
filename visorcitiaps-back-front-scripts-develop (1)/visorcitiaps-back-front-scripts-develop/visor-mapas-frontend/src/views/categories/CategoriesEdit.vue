<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Editar categoría</h2>
    </div>

    <div class="container-action">
        
    </div>

    <div class="container-body">
        <label for="name">Nombre</label>
        <input id="name" v-model="category.name" type="text"
        name="nombre" :class="{'error' : errors.has('nombre') }" v-validate="'required'"
        :placeholder="[[errors.has('nombre') ? errors.first('nombre') : 'Ingresa el nombre de la categoría']]"/>

        <label for="desc">Descripción</label>
        <textarea id="desc" v-model="category.desc" placeholder="Ingresa la descripción de la categoría"></textarea>


        <div class="container-action-left" style="justify-content: start;">
            <button class="btn btn-secondary" v-on:click="onClickGoBack()">Volver</button>
            <button class="btn btn-secondary" v-on:click="onClickGoNext()">Actualizar categoría</button>
        </div>

    </div>
 
</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'

import {restCategory} from '@/components/mixin/restCategory.js'

export default{
    name: 'test',
    mixins: [restCategory, restCategory],
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
                {name: 'Capas', path: '/capas'},
                {name: 'Categorías', path: '/capas/categorias-de-capas/'},
                {name: this.category.name, path: '/capas/categorias-de-capas/' + this.$route.params.idCategory},
                {name: 'Editar categoría'}
          ]
        }
    },
    props: [],
    methods:{
        onClickGoBack: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/capas/categorias-de-capas/' + this.$route.params.idCategory)
        },
        onClickGoNext: function() {
            var that = this;
            this.$validator.validateAll()
            .then(function (response) {
                if(response) {
                    that.putCategory();
                } else {
                    console.log('Error')
                }
            })
        }
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.layers? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getCategory();
    }
}
</script>

<style lang="scss" scoped>

</style>
