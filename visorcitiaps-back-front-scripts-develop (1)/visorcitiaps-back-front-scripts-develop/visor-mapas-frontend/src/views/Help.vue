<template>
<main class="container">
    <h2>¿Necesitas ayuda?</h2>
    <span>En esta guía puedes encontrar información sobre los controles del panel de administración</span>

    <h3>Navegación por el sitio</h3>
    <span>Visita las diferentes secciones del sitio utilizando la barra superior</span>
    <br>
    <img src="/images/tutorial/navigation.png"/>

    <h3>Lista de recursos</h3>
    <span>Consulta las listas de elementos cambiando las páginas.</span>
    <br>
    <span>Se puede seleccionar un recurso para ver su información completa.</span>
    <br>
    <span>Las opciones en el lado derecho permiten agregar un nuevo elemento y ver otros recursos relacionados.</span>
    <br>
    <img src="/images/tutorial/pagination.png"/>

    <h3>Información de recurso</h3>
    <span>Visualiza la información completa de un recurso.</span>
    <br>
    <span>Las opciones en el lado derecho permiten realizar cambios y acciones relacionadas con el recurso.</span>
    <br>
    <img src="/images/tutorial/item.png"/>

    <h3>Asignación de capas</h3>
    <span>Asigna una capa a los usuarios deseados.</span>
    <br>
    <span>Las asignaciones pueden cancelarse en cualquier momento.</span>
    <br>
    <img src="/images/tutorial/assign.png"/>

    <h3>Dispositivos móviles</h3>
    <span>¿Estás en tu smartphone o tablet? Accede al panel de administración y haz uso de sus funciones gracias a su diseño adaptable</span>
    <br>
    <img style="max-width: 40%;" src="/images/tutorial/mobile.png"/>
    
    <h3>Atajos de teclado</h3>

    <span>Para hacer uso de las acciones del lado derecho, presionar la tecla ctrl junto a la tecla correspondiente a la primera letra del botón.</span>
    <br>
    <span>Por ejemplo, si la acción es Editar Capa</span>
    <br>
    <img src="/images/tutorial/option.png"/>
    <br>
    <span>La combinación de teclas es la siguiente:</span>

    <div class="button-explanation separation">
        <div class="key">
            <span style="font-size: 0.8em;">ctrl</span>
        </div>
        <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
        <div class="key">
            <span>E</span>
        </div>
    </div>

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

    },
    updated () {

    },
    created () {
        //Verificar permisos
        this.$store.getters.getPermissions? true:this.$router.push('/inicio')
    },
    mounted () {
    }
}
</script>

<style lang="scss" scoped>
button{
    margin-top: 1em;
}
.ol-control {
    position: relative;
}
div{
    width: min-content;
}
.ol-control button:hover {
    background-color: #036016;
    color: white;
}
p {
    margin-left: 1em;
}
h3{
    margin-bottom: 0;
}
.button-explanation{
    display: flex;
    flex-direction: row;
    width: 100%;
    align-items: center;
}
.separation{
    margin-top: 0.5em;
    padding-top: 0.5em;
}
.border {
    border-top: solid 1px gray;
}
.key{
    background-color: gray;
    width: 2em;
    height: 2em;
    text-align: center;
    color: white;
    border-radius: 4px;
    font-size: 1.25em;
    padding: 0.25em;
    font-weight: bold;
}
img {
    max-width: 100%;
    margin-top: 0.5em;
}

</style>
