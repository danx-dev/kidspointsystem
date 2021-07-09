<template>
    <div>
        <b-row class="pb-5">
        <h1>{{name}}</h1>
        </b-row>
        <assignment-table title="Open Assignments" :assignments="openAssignments" :fields="fields" />
        <assignment-table title="Closed Assignments" :assignments="closedAssignments" class="pt-5" :fields="fields" />
    </div>
</template>


<script>
import AssignmentTable from '../components/AssignmentTable.vue'
import axios from 'axios'
import config from '../config'
export default {
  components: { AssignmentTable },
    data(){
        return {
            name: '',
            fields: ["Id", "Description", "Points"],
            allAssignments: []
        }
    },
    created: async function(){
        var self = this;
        await axios.get("/kidsinfo?kidid=" + config.kidId).then(resp => {
            console.log('KidInfo',resp.data)
            self.name = resp.data.Name
        })
        await axios.get("/kidsassignments?kidid=" + config.kidId ).then(resp => {
            console.log('Assignments',resp.data)
            self.allAssignments = resp.data;
        })
    },
    computed: {
        openAssignments: function(){
            return this.allAssignments.filter(el => el.Status === 'Open')
        },
        closedAssignments: function(){
            return this.allAssignments.filter(el => el.Status === 'Closed')
        }
    }
}
</script>