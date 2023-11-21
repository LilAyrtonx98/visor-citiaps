<template>
<main class="container">
    <div>
        <h1>Visor de Mapas</h1>
        <p>Sebastián Pinto-Agüero</p>
        <p>2019-2020</p>
        <p>César Salazar Silva</p>
        <p>2021-2022</p>

        <br>
        <b>Desarrollado gracias a:</b>

        <ul>
            <li><a href="https://github.com/openlayers/openlayers">OpenLayers</a></li>
            <li><a href="https://github.com/vuejs/vue">Vue.js</a></li>
            <li><a href="https://github.com/apexcharts/vue-apexcharts">ApexCharts.js</a></li>
        </ul>        
    </div>
    <div>
        <p>Imagen de fondo por Jorge Barahona en Unsplash</p>
        <a style="background-color:black;color:white;text-decoration:none;padding:4px 6px;font-size:0.8em;font-weight:bold;line-height:1.2;display:inline-block;border-radius:3px" href="https://unsplash.com/@jorgebarahona?utm_medium=referral&amp;utm_campaign=photographer-credit&amp;utm_content=creditBadge" target="_blank" rel="noopener noreferrer" title="Download free do whatever you want high-resolution photos from Jorge Barahona">
            <span style="display:inline-block;padding:2px 3px">
                <svg
                    xmlns="http://www.w3.org/2000/svg" style="height:12px;width:auto;position:relative;vertical-align:middle;top:-2px;fill:white" viewBox="0 0 32 32">
                    <title>unsplash-logo</title>
                    <path d="M10 9V0h12v9H10zm12 5h10v18H0V14h10v9h12v-9z"></path>
                </svg>
            </span>
            <span style="display:inline-block;padding:2px 3px">Jorge Barahona</span>
        </a>
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
        hello: function() {
        console.log("Hello");
        }
    },
    updated () {

    },
    created () {

    },
    mounted () {

    }
}
</script>

<style lang="scss" scoped>
.container {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}
</style>