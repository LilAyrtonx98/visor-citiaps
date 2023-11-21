<template>
<main class="container">
    <h2>Esta página no existe</h2>
    <p>No se ha podido encontrar lo que estás buscando.</p>
    <p>Esto se debe a que la página ya no existe o no se encuentra disponible en este momento.</p>
    <button class="btn btn-secondary" v-on:click="onClickGoBack()">Volver</button>
</main>
</template>

<style>

</style>

<script>
// @ is an alias to /src

export default{
    name: 'test',
    data(){
        return{
        
        }
    },
    components:{
        
    },
    computed: {

    },
    props: [],
    methods:{
        hello: function() {
            console.log("What are you looking for?");
        },
        onClickGoBack: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/inicio')
        },
    },
    updated () {

    },
    created () {

    },
    mounted () {
        this.hello();
    }
}
</script>

<style lang="scss" scoped>
button{
    margin-top: 1em;
}
</style>
