import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import NotFound from './views/NotFound.vue'
import Dashboard from './views/Dashboard.vue'
import Logout from './views/Logout.vue'
import About from './views/About.vue'
import Help from './views/Help.vue'

// Visor
import VisorViewer from './views/visor/VisorViewer.vue'
import VisorMaps from './views/visor/VisorMaps.vue'
import VisorHelp from './views/visor/VisorHelp.vue'

// Mapa
import MapViewer from './views/maps/MapViewer.vue'
import Maps from './views/maps/Maps.vue'
import MapsInfo from './views/maps/MapsInfo.vue'
import MapsNew from './views/maps/MapsNew.vue'
import MapsEdit from './views/maps/MapsEdit.vue'
import MapsLogs from './views/maps/MapsLogs.vue'
import MapsDelete from './views/maps/MapsDelete.vue'
import MapsAssignLayers from './views/maps/MapsAssignLayers.vue'
import MapsAssignUsers from './views/maps/MapsAssignUsers.vue'
import MapsAssignGeoprocessings from './views/maps/MapsAssignGeoprocessings.vue'

// Capas
import Layers from './views/layers/Layers.vue'
import LayersInfo from './views/layers/LayersInfo.vue'
import LayersNew from './views/layers/LayersNew.vue'
import LayersEdit from './views/layers/LayersEdit.vue'
import LayersLogs from './views/layers/LayersLogs.vue'
import LayersDelete from './views/layers/LayersDelete.vue'
import LayersAssignMaps from './views/layers/LayersAssignMaps.vue'

// Categorías
import Categories from './views/categories/Categories.vue'
import CategoriesInfo from './views/categories/CategoriesInfo.vue'
import CategoriesNew from './views/categories/CategoriesNew.vue'
import CategoriesEdit from './views/categories/CategoriesEdit.vue'
import CategoriesLogs from './views/categories/CategoriesLogs.vue'
import CategoriesDelete from './views/categories/CategoriesDelete.vue'

// Geoprocesos
import Geoprocessings from './views/geoprocessings/Geoprocessings.vue'
import GeoprocessingsInfo from './views/geoprocessings/GeoprocessingsInfo.vue'
import GeoprocessingsNew from './views/geoprocessings/GeoprocessingsNew.vue'
import GeoprocessingsEdit from './views/geoprocessings/GeoprocessingsEdit.vue'
import GeoprocessingsLogs from './views/geoprocessings/GeoprocessingsLogs.vue'
import GeoprocessingsDelete from './views/geoprocessings/GeoprocessingsDelete.vue'
import GeoprocessingsAssignMaps from './views/geoprocessings/GeoprocessingsAssignMaps.vue'

// Usuarios
import Users from './views/users/Users.vue'
import UsersInfo from './views/users/UsersInfo.vue'
import UsersNew from './views/users/UsersNew.vue'
import UsersEdit from './views/users/UsersEdit.vue'
import UsersLogs from './views/users/UsersLogs.vue'
import UsersDelete from './views/users/UsersDelete.vue'
import UsersAssignMaps from './views/users/UsersAssignMaps.vue'

//Grupos
import Groups from './views/groups/Groups.vue'
import GroupsInfo from './views/groups/GroupsInfo.vue'
import GroupsNew from './views/groups/GroupsNew.vue'
import GroupsEdit from './views/groups/GroupsEdit.vue'
import GroupsLogs from './views/groups/GroupsLogs.vue'
import GroupsDelete from './views/groups/GroupsDelete.vue'

// Cuenta de usuario
import Account from './views/account/Account.vue'
import AccountChangePassword from './views/account/AccountChangePassword.vue'

const originalPush = Router.prototype.push;
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
};

Vue.use(Router)

export default new Router({
  //mode: 'history',
  routes: [
    { path: '/', redirect: '/inicio' },
    {
      path: '/inicio',
      name: 'inicio',
      component: Home
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: Dashboard
    },
    {
      path: '/visor/ayuda',
      name: 'ayuda del visor',
      component: VisorHelp
    },
    {
      path: '/visor',
      name: 'mapas visor',
      component: VisorMaps
    },
    {
      path: '/visor/:idMap',
      name: 'visor de este mapa',
      component: VisorViewer
    },
    {
      path: '/acerca-de',
      name: 'acerca de',
      component: About
    },
    {
      path: '/ayuda',
      name: 'ayuda',
      component: Help
    },
    {
      path: '/mapas/:idMap/visor',
      name: 'visor',
      component: MapViewer
    },
    {
      path: '/mapas',
      name: 'mapas',
      component: Maps
    },
    {
      path: '/mapas/nuevo-mapa',
      name: 'nuevo-mapa',
      component: MapsNew
    },
    {
      path: '/mapas/:idMap',
      name: 'mapa-idmapa',
      component: MapsInfo
    },
    {
      path: '/mapas/:idMap/editar-mapa',
      name: 'editar-mapa',
      component: MapsEdit
    },
    {
      path: '/mapas/:idMap/historial-de-cambios',
      name: 'historial-de-cambios-mapa',
      component: MapsLogs
    },
    {
      path: '/mapas/:idMap/eliminar-mapa',
      name: 'eliminar-mapa',
      component: MapsDelete
    },
    {
      path: '/mapas/:idMap/asignar-capas',
      name: 'asignar-capas-mapa',
      component: MapsAssignLayers
    },
    {
      path: '/mapas/:idMap/asignar-usuarios',
      name: 'asignar-usuarios-mapa',
      component: MapsAssignUsers
    },
    {
      path: '/mapas/:idMap/asignar-geoprocesos',
      name: 'asignar-geoprocesos-mapa',
      component: MapsAssignGeoprocessings
    },
    {
      path: '/cuenta/cerrar-sesion',
      name: 'cerrar sesion',
      component: Logout
    },
    {
      path: '/capas',
      name: 'capas',
      component: Layers
    },
    {
      path: '/capas/categorias-de-capas',
      name: 'categorias',
      component: Categories
    },
    {
      path: '/capas/categorias-de-capas/nueva-categoria',
      name: 'nueva-categoria',
      component: CategoriesNew
    },
    {
      path: '/capas/nueva-capa',
      name: 'nueva-capa',
      component: LayersNew
    },
    {
      path: '/capas/:idLayer',
      name: 'capas-idlayer',
      component: LayersInfo
    },
    {
      path: '/capas/:idLayer/editar-capa',
      name: 'editar-capa',
      component: LayersEdit
    },
    {
      path: '/capas/:idLayer/asignar-mapas',
      name: 'capa-asignar-mapas',
      component: LayersAssignMaps
    },
    {
      path: '/capas/:idLayer/historial-de-cambios',
      name: 'historial-de-cambios',
      component: LayersLogs
    },
    {
      path: '/capas/:idLayer/eliminar-capa',
      name: 'eliminar-capa',
      component: LayersDelete
    },
    {
      path: '/capas/categorias-de-capas/:idCategory',
      name: 'categorias-idcategory',
      component: CategoriesInfo
    },
    {
      path: '/capas/categorias-de-capas/:idCategory/editar-categoria',
      name: 'editar-categoria',
      component: CategoriesEdit
    },
    {
      path: '/capas/categorias-de-capas/:idCategory/historial-de-cambios',
      name: 'historial-de-cambios-categoria',
      component: CategoriesLogs
    },
    {
      path: '/capas/categorias-de-capas/:idCategory/eliminar-categoria',
      name: 'eliminar-categoria',
      component: CategoriesDelete
    },
    {
      path: '/geoprocesos',
      name: 'geoprocesos',
      component: Geoprocessings
    },
    {
      path: '/geoprocesos/nuevo-geoproceso',
      name: 'nuevo-geoproceso',
      component: GeoprocessingsNew
    },
    {
      path: '/geoprocesos/:idGeo',
      name: 'geoprocesos-idgeo',
      component: GeoprocessingsInfo
    },
    {
      path: '/geoprocesos/:idGeo/asignar-mapas',
      name: 'geoproceso-asignar-mapas',
      component: GeoprocessingsAssignMaps
    },
    {
      path: '/geoprocesos/:idGeo/editar-geoproceso',
      name: 'editar-geoproceso',
      component: GeoprocessingsEdit
    },
    {
      path: '/geoprocesos/:idGeo/historial-de-cambios',
      name: 'historial-de-cambios-geo',
      component: GeoprocessingsLogs
    },
    {
      path: '/geoprocesos/:idGeo/eliminar-geoproceso',
      name: 'eliminar-geoproceso',
      component: GeoprocessingsDelete
    },
    {
      path: '/usuarios',
      name: 'usuarios',
      component: Users
    },
    {
      path: '/usuarios/grupos-de-usuarios',
      name: 'grupos',
      component: Groups
    },
    {
      path: '/usuarios/grupos-de-usuarios/nuevo-grupo',
      name: 'nuevo-grupo',
      component: GroupsNew
    },
    {
      path: '/usuarios/grupos-de-usuarios/:idGroup',
      name: 'grupos-idgroup',
      component: GroupsInfo
    },
    {
      path: '/usuarios/grupos-de-usuarios/:idGroup/editar-grupo',
      name: 'edit-grupo',
      component: GroupsEdit
    },
    {
      path: '/usuarios/grupos-de-usuarios/:idGroup/historial-de-cambios',
      name: 'historial-de-cambios-grupo',
      component: GroupsLogs
    },
    {
      path: '/usuarios/grupos-de-usuarios/:idGroup/eliminar-grupo',
      name: 'eliminar-grupo',
      component: GroupsDelete
    },
    {
      path: '/usuarios/nuevo-usuario',
      name: 'nuevo-usuario',
      component: UsersNew
    },
    {
      path: '/usuarios/:idUser',
      name: 'usuarios-iduser',
      component: UsersInfo
    },
    {
      path: '/usuarios/:idUser/editar-usuario',
      name: 'edit-usuario',
      component: UsersEdit
    },
    {
      path: '/usuarios/:idUser/historial-de-cambios',
      name: 'historial-de-cambios-usuario',
      component: UsersLogs
    },
    {
      path: '/usuarios/:idUser/eliminar-usuario',
      name: 'eliminar-usuario',
      component: UsersDelete
    },
    {
      path: '/usuarios/:idUser/asignar-mapas',
      name: 'asignar-mapas',
      component: UsersAssignMaps
    },
    { 
      path: "/cuenta/mi-cuenta",
      name: 'cuenta usuario',
      component: Account 
    },
    { 
      path: "/cuenta/cambiar-contrasena",
      name: 'cambiar contrasena',
      component: AccountChangePassword 
    },
    { path: '*', redirect: '/pagina-no-encontrada' },
    { 
      path: "/pagina-no-encontrada",
      name: 'not-found',
      component: NotFound 
    }
  ],
  linkActiveClass: "active", // active class for non-exact links.
  //linkExactActiveClass: "active" // active class for *exact* links.
})
