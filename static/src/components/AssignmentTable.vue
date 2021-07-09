<template>
    <b-row>
        <b-row>
            <h3>{{title}}</h3>
        </b-row>
        <b-row>
            <b-table striped hover :items="assignments" :fields="fields">
                <template #cell(actions)="data" v-if="showAction">
                    <b-button variant="success" :data-id="data.item.id">Take Me</b-button>
                </template>
            </b-table>
        </b-row>
    </b-row>
</template>

<script>
export default {
    props:{
        title: {
            required: true, type: String},
        assignments: {required: true, type: Array},
        showAction: {
            type: Boolean,
            default: true
        },
        fields: {required:true, type: Array}
    },
    data(){
        return {
        }
    },
    created: function(){
        console.log('Table Created', this.showAction)
    }
}
</script>