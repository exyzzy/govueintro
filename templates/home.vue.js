{{ define "vue" }}
 <script>
new Vue({
    delimiters: ["${", "}"], //required to not conflict with go template action delimiters
    el: "#app",
    vuetify: new Vuetify(),
    data: () => ({})
});
</script>
{{ end }}