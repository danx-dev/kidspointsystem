<template>
  <div>
    <b-table striped hover :items="items"></b-table>
  </div>
</template>

<script>
import axios from 'axios';
  export default {
    data() {
      return {
        items: [
          { id: 1, name: 'Ella', points: 0 },
          { id: 2, name: 'Frida', points: 0 },
          { id: 3, name: 'Karla', points: 0 },
        ]
      }
    },
    created: async function(){
      var self = this;
      await axios.get("/kids")
      .then(resp => {
        console.log(resp)
        self.items = resp.data;
      })
    }
  }
</script>