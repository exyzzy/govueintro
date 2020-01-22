Vue.component('todo-list', {
    // delimiters: ['${', '}'],   CUSTOM DELIMITERS DO NOT WORK IN COMPONENTS, AND ARE NOT NEEDED BY GO!!
    data: () => ({
        todos: []
    }),
    mounted: function() {
        this.$nextTick(function () {
            this.loadTodos();
        })
    },
    methods: {
        addTodo() {

            axios.post('/api/todo', {"title": "Todo"}
            )
            .then((response) => {
                this.todos.unshift(response.data)
            })
            .catch((error) => {
                if (error.response) {
                // The request was made and the server responded with a status code
                // that falls out of the range of 2xx
                console.log(error.response.data);
                console.log(error.response.status);
                console.log(error.response.headers);
                } else if (error.request) {
                // The request was made but no response was received
                // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
                // http.ClientRequest in node.js
                console.log(error.request);
                } else {
                // Something happened in setting up the request that triggered an Error
                console.log('Error', error.message);
                }
                console.log(error.config);
            });

        },
        deleteTodo(item) {
            axios.delete('/api/todo/'+item.id.toString()
            )
            .then((response) => {
                let index = this.todos.indexOf(item)
                this.todos.splice(index, 1)
            })
            .catch((error) => {
                if (error.response) {
                console.log(error.response.data);
                console.log(error.response.status);
                console.log(error.response.headers);
                } else if (error.request) {
                console.log(error.request);
                } else {
                console.log('Error', error.message);
                }
                console.log(error.config);
            });
            
        },
        updateTodo(item) {
            axios.put('/api/todo/'+item.id.toString(), 
                item 
            )
            .then((response) => {
                //vuetify handles checkbox/textfiled update
                return
            })
            .catch((error) => {
                if (error.response) {
                console.log(error.response.data);
                console.log(error.response.status);
                console.log(error.response.headers);
                } else if (error.request) {
                console.log(error.request);
                } else {
                console.log('Error', error.message);
                }
                console.log(error.config);
            });

        },
        loadTodos () {
            axios.get('/api/todos', 
            )
            .then((response) => {

                if (response.data == null ) {
                    console.log("no todos returned")
                } else {

                    for (let i= 0; i < response.data.length; i++) {
                        this.todos.push(response.data[i])
                    }

                }
            })
            .catch((error) => {
                if (error.response) {
                    console.log(error.response.data);
                    console.log(error.response.status);
                    console.log(error.response.headers);
                    } else if (error.request) {
                    console.log(error.request);
                    } else {
                    console.log('Error', error.message);
                    }
                    console.log(error.config);
                });
        }

    },
    template: `
    <v-card
    max-width="600"
    class="mx-auto"
    >
        <v-toolbar
            color="green accent-1"
        >
            <v-toolbar-title>My Todos</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn
                color="blue darken-1"
                @click="addTodo()"
                dark
            >
            Add Todo
            </v-btn>
        </v-toolbar>
        <v-list >
            <v-list-item 
                v-for="todo in todos"
                :key = todo.id
            >
                    <template v-slot:default="{ active, toggle }">
                    <v-list-item-action>
                        <v-checkbox
                        v-model="todo.done"
                        @change="updateTodo(todo)"
                        color="blue darken-1"
                        ></v-checkbox>
                    </v-list-item-action>
    
                    <v-list-item-content>
                        <v-list-item-title>
                            <v-text-field 
                                dense 
                                hide-details 
                                v-model="todo.title"
                                @change="updateTodo(todo)"
                            ></v-text-field>
                        </v-list-item-title>
                    </v-list-item-content>
                    <v-list-item-action>
                        <v-tooltip top>
                            <template v-slot:activator="{ on }">
                                <v-icon 
                                @click="deleteTodo(todo)"
                                v-on="on"
                                color="blue darken-1"
                                >
                                mdi-delete
                                </v-icon>
                            </template>
                            Delete Todo
                        </v-tooltip>            
                    </v-list-item-action>
                    </template>
            </v-list-item>
        </v-list>
    </v-card>
    `
});

