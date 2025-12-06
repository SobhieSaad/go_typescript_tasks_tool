import { Container, Stack } from "@chakra-ui/react";
import Navbar from "./components/Navbar";
import TodoForm from "./components/TodoForm";
import TodoList from "./components/TodoList";

export const BASE_URL = "/api";
//for development, use "http://localhost:5000/api" or the port you set in your backend

function App() {
	return (
		<Stack h='100vh'>
			<Navbar />
			<Container>
				<TodoForm />
				<TodoList />
			</Container>
		</Stack>
	);
}

export default App;
