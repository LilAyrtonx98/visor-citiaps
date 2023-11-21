<template>
<main class="container">
    <h2>Ayuda para usar el visor de mapas</h2>
    <span>En esta guía puedes encontrar información sobre los controles del visor de mapas</span>

    <h3>Controles del mapa</h3>

    <span>Interacciones principales relacionadas con el mapa</span>

    <div class="button-explanation">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <button class="icon zoom-in"></button>
        </div>
        <p>Aumenta el zoom al mapa</p>
    </div>

    <div class="button-explanation">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <button class="icon zoom-out"></button>
        </div>
        <p>Reduce el zoom al mapa</p>
    </div>

    <div class="button-explanation">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <button class="icon legend"></button>
        </div>
        <p>Muestra la simbología de las capas en el mapa</p>
    </div>

    <div class="button-explanation">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <button class="icon geolocation"></button>
        </div>
        <p>Muestra la ubicación del usuario en el mapa</p>
    </div>
    
    <div class="button-explanation">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <button class="icon full-screen"></button>
        </div>
        <p>Carga el mapa en toda la pantalla</p>
    </div>

    <h3>Controles adicionales</h3>

    <span>Interacciones secundarias</span>

    <div class="button-explanation">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <button class="icon toolbox"></button>
        </div>
        <p>Abre la caja de herramientas</p>
    </div>

    <h3>Caja de herramientas</h3>

    <span>Herramientas que permiten realizar acciones en el mapa</span>

    <div class="button-explanation separation" style="color: #036016;">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <font-awesome-icon icon="hand-pointer"/>
        </div>
        <strong style="margin-left: 1em;">Seleccionar punto</strong>
    </div>
    <span>Permite hacer click en el mapa y ver la información sobre las capas que se encuentran en el punto seleccionado</span>

    <div class="button-explanation separation border" style="color: #036016;">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <font-awesome-icon icon="layer-group"/>
        </div>
        <strong style="margin-left: 1em;">Capas</strong>
    </div>
    <span>Despliega una lista de capas que pueden seleccionarse para ser mostradas en el mapa</span>

    <div class="button-explanation separation border" style="color: #036016;">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <font-awesome-icon icon="chart-bar"/>
        </div>
        <strong style="margin-left: 1em;">Gráfico</strong>
    </div>
    <span>Muestra gráficos con datos de acuerdo a la capa y al área seleccionada por el usuario</span>

    <div class="button-explanation separation border" style="color: #036016;">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <font-awesome-icon icon="ruler-combined"/>
        </div>
        <strong style="margin-left: 1em;">Medir mapa</strong>
    </div>
    <span>Calcula el área de una zona o la distancia entre una secuencia de puntos</span>

    <div class="button-explanation separation border" style="color: #036016;">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <font-awesome-icon icon="search"/>
        </div>
        <strong style="margin-left: 1em;">Buscar ubicación</strong>
    </div>
    <span>Muestra en el mapa la ubicación buscada por el usuario</span>

    <div class="button-explanation separation border" style="color: #036016;">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <font-awesome-icon icon="file-export"/>
        </div>
        <strong style="margin-left: 1em;">Exportar mapa</strong>
    </div>
    <span>Guarda el mapa como imagen PNG o documento PDF</span>

    <div class="button-explanation separation border" style="color: #036016;">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <font-awesome-icon icon="file-import"/>
        </div>
        <strong style="margin-left: 1em;">Subir capa</strong>
    </div>
    <span>Sube una capa al servidor de mapas</span>

    <div class="button-explanation separation border" style="color: #036016;">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <font-awesome-icon icon="map"/>
        </div>
        <strong style="margin-left: 1em;">Mapa base</strong>
    </div>
    <span>Cambia el mapa que se muestra</span>

    <div class="button-explanation separation border" style="color: #036016;">
        <div class="custom-control-no-position ol-unselectable ol-control">
            <font-awesome-icon icon="times"/>
        </div>
        <strong style="margin-left: 1em;">Salir</strong>
    </div>
    <span>Vuelve a la página de inicio</span>

    <h3>Atajos de teclado</h3>

    <span>Realiza acciones mediante el teclado</span>

    <div class="button-explanation separation">
        <div class="key">
            ctrl
        </div>
        <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
        <div class="key">
            A
        </div>
        <p>Muestra los atajos del teclado en el mapa</p>
    </div>

    <div class="button-explanation separation">
        <div class="key">
            <span style="font-size: 0.85em;">shift</span>
        </div>
        <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
        <div class="key">
            <font-awesome-icon icon="arrow-up"/>
        </div>
        <p>Aumenta el zoom al mapa</p>
    </div>

    <div class="button-explanation separation">
        <div class="key">
            <span style="font-size: 0.85em;">shift</span>
        </div>
        <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
        <div class="key">
            <font-awesome-icon icon="arrow-down"/>
        </div>
        <p>Reduce el zoom al mapa</p>
    </div>

    <div class="button-explanation separation">
        <div class="key">
            ctrl
        </div>
        <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
        <div class="key">
            S
        </div>
        <p>Muestra la simbología de las capas en el mapa</p>
    </div>

    <div class="button-explanation separation">
        <div class="key">
            ctrl
        </div>
        <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
        <div class="key">
            U
        </div>
        <p>Muestra la ubicación del usuario en el mapa</p>
    </div>

    <div class="button-explanation separation">
        <div class="key">
            ctrl
        </div>
        <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
        <div class="key">
            P
        </div>
        <p>Carga el mapa en toda la pantalla</p>
    </div>

    <div class="button-explanation separation">
        <div class="key">
            ctrl
        </div>
        <span style="margin: 0 0.5em; font-weight: bold; font-size: 1.5em;">+</span>
        <div class="key">
            H
        </div>
        <p>Abre la caja de herramientas</p>
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
        this.$store.getters.getPermissions.visor? true:this.$router.push('/pagina-no-encontrada')
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

</style>
