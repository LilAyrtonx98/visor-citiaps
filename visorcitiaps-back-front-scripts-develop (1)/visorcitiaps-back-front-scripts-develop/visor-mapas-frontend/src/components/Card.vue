<template>
<router-link style="text-decoration: none;" :to="url">
  <div class="card">
      <img :src="imgsrc" :alt="imgalt"/>
      <div class="content">
          <p>{{title}}</p>
          <div>{{subtitle}}</div>
      </div>
  </div>
</router-link>
</template>

<script>
export default{
  props:['imgsrc', 'title', 'subtitle', 'imgalt', 'url'],
  data(){
    return {
        
    }
  },
  computed: {

  },
  methods: {

  },
  created(){
  
  },
  watch: {

  }
}
</script>

<style lang="scss" scoped>

.card {
    display: grid;
    box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
    transition: 0.3s;
    height: 25vh;
    img {
        justify-self: center;
        object-fit: cover;
        width: 100%;
        height: 100%;
    }
    .content {
        padding: 2px 16px;
        margin: 1em 0;
        p {
            font-weight: bold;
        }
    }
}

.card:hover {
  box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2);
}

a {
  color: black; 
  font-weight: unset;
}

a:hover {
  color: rgba(3, 96, 22, 0.9);
}
</style>