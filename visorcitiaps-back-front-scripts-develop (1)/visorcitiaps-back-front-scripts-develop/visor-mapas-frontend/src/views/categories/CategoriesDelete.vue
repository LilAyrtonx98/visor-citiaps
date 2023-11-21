<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>{{category.name}}</h2>
    </div>

    <div class="container-body">

        <span>¿Realmente deseas eliminar esta categoría?</span>
        <br>
        <span>Si lo haces, un registro de este recurso quedará en la base de datos</span>
        <br><br>
        <button class="btn btn-danger" style="margin-right: 1em;" v-on:click="deleteCategory()">Sí</button>
        <button class="btn btn-secondary" v-on:click="onClickDontDelete()">No</button>

    </div>

</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'

import {restCategory} from '@/components/mixin/restCategory.js'

export default{
    name: 'test',
    mixins: [restCategory],
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
                {name: 'Capas', path: '/capas'},
                {name: 'Categorías', path: '/capas/categorias-de-capas/'},
                {name: this.category.name, path: '/capas/categorias-de-capas/' + this.$route.params.idCategory},
                {name: 'Eliminar categoría'}
          ]
        }
    },
    props: [],
    methods:{
        onClickDontDelete: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/capas/categorias-de-capas/' + this.$route.params.idCategory)
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
