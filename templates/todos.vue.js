{{ define "vue" }}
<script src="/static/component/vueTodo.js"></script>
<script>
new Vue({
    delimiters: ["${", "}"],
    el: "#app",
    vuetify: new Vuetify(),
    data: () => ({})
});
</script>
{{ end }}