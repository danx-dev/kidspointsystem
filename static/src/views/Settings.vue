<template>
  <div>
      <h1>Settings</h1>

    <b-form-group
      id="fieldset-1"
      label="Enter your App Token"
      label-for="input-1"
      :invalid-feedback="invalidFeedback"
      :state="tokenState"
    >
      <b-form-input id="input-1" v-model="token" :state="tokenState" trim></b-form-input>
    </b-form-group>
    <br/><br/><br/>
    <b-form-group
      id="fieldset-1"
      label="Enter your Kid Id"
      label-for="input-1"
      :invalid-feedback="invalidFeedback"
    >
      <b-form-input id="input-1" v-model="kidId" type="number"></b-form-input>
    </b-form-group>
    <br/><br/><br/>
    <b-form-group>
        <b-button variant="success" @click="save">Save</b-button>
    </b-form-group>
  </div>
</template>

<script>
  export default {
    computed: {
      tokenState() {
        return this.token.length >= 4
      },
      invalidFeedback() {
        if (this.token.length > 0) {
          return 'Enter at least 4 characters.'
        }
        return 'Please enter something.'
      }
    },
    data() {
      return {
        token: '',
        kidId: 0
      }
    },
    created: function(){
      this.AppToken = localStorage.AppToken;
      this.KidId = localStorage.KidId;
    },
    methods: {
        save: function(){
            localStorage.AppToken = this.token;
            localStorage.KidId = this.kidId;
        }
    }
  }
</script>