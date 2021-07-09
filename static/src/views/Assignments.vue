<template>
  <div>
    <b-table striped hover :items="items" :fields="fields">
        <template #cell(actions)="data">
            <b-button variant="success" :data-id="data.item.id">Take Me</b-button>
      </template>
    </b-table>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        fields: [
            "id",
            "name",
            "points",
            "actions"
        ],
        items: [
          { id: 1, name: 'Tømme opvasker', points: 5 },
          { id: 2, name: 'Parre Strømper', points: 5 },
          { id: 3, name: 'Rydde op i stuen', points: 5 },
        ]
      }
    }
  }
</script>