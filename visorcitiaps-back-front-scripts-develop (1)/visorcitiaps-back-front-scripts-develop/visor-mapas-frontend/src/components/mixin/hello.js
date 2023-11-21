export const hello = {
    data() {
        return {
            message: 'Hola'
        }
    },
    methods: {
      sayHello: function () {
        console.log(this.message + ' desde mixin!')
      }
    }
  }