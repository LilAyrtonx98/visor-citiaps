<template>
  <body id="app">

    <HeaderCollapsible
    id="header"
    v-bind:links="routes"
    v-bind:account="isAccountVisible"
    >
    </HeaderCollapsible>

    <router-view/>

    <FooterFixed
    id="footer"
    >
    </FooterFixed>
     

  </body>
</template>

<script>
import HeaderCollapsible from '@/components/HeaderCollapsible.vue'
import FooterFixed from '@/components/FooterFixed.vue'

import {vuex} from '@/components/mixin/vuex.js'

export default{
  name: 'home',
  data(){
    return{

    }
  },
  components:{
      HeaderCollapsible,
      FooterFixed
  },
  mixins: [
    vuex
    ],
  computed: {
    isAccountVisible: function() {
      if(this.$store.getters.getPermissions) {
        return true;
      }
      return false;
    },
    routes: function(){
      let permissions = this.$store.getters.getPermissions;
      let routesByPermissions = [];
      
      //Si no ha iniciado sesión se muestran las rutas básicas
      if(!permissions) {
        routesByPermissions.push({name:'Inicio' , path: '/inicio'});
        routesByPermissions.push({name:'Acerca de' , path: '/acerca-de'});
        //routesByPermissions.push({name:'Iniciar sesión' , path: '/iniciar-sesion'});
      } else {
        routesByPermissions.push({name:'Inicio' , path: '/dashboard'});
        if(permissions.visor) {
          routesByPermissions.push({name:'Visor' , path: '/visor'});
        }
        if(permissions.maps) {
          routesByPermissions.push({name:'Mapas' , path: '/mapas'});
        }
        if(permissions.layers) {
          routesByPermissions.push({name:'Capas' , path: '/capas'});
        }
        if(permissions.geo) {
          routesByPermissions.push({name:'Geoprocesos' , path: '/geoprocesos'});
        }
        if(permissions.users) {
          routesByPermissions.push({name:'Usuarios' , path: '/usuarios'});
        }
      }
      return routesByPermissions;
    },
  },
  props: [],
  methods:{

  },
  beforeUpdate () {
    if(this.$route.fullPath != '/acerca-de') {
      this.$store.getters.getPermissions? true:this.$router.push('/inicio')
    }
  },
  updated () {
    
  },
  created () {
    //Si se está accediendo al sitio en una nueva pestaña, intenta recuperar la sesión pública
    //Verifica si existe la cookie session_public y escribe el rol del usuario en Vuex (válido solo para esta pestaña)
    //Lo mismo para token de sesión
    this.$store.commit('setPermissions');
    //this.$store.commit('setToken');

    //Por defecto, se envía el token de sesión en cada consulta
    //this.$http.defaults.headers.common = {'Authorization': 'Bearer ' + this.$store.getters.getToken}
  },
  mounted () {
    if(this.$route.fullPath != '/acerca-de') {
      this.$store.getters.getPermissions? true:this.$router.push('/inicio')
    }
    if(window.screen.width > 640) {
      this.$store.commit('setIsNotMobile');
    } else {
      this.$store.commit('setIsMobile');
    }
  },
  watch: {
    'isMapActive': function (newValue, oldValue) {
      if(newValue) {
        let footer = document.getElementById("footer");
        footer.style.display = 'none';
        console.log('hide footer')
      } else {
        let footer = document.getElementById("footer");
        footer.style.display = 'grid';
        console.log('show footer')
      }
    }
  }
}
</script>