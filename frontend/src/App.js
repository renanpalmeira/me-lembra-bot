import ChatBot from 'react-simple-chatbot'
import axios from 'axios'
import validator from "validar-telefone"
import { ThemeProvider } from 'styled-components'

// all available props
const theme = {
  background: '#f5f8fb',
  fontFamily: 'Helvetica Neue',
  headerBgColor: '#EF6C00',
  headerFontColor: '#fff',
  headerFontSize: '15px',
  botBubbleColor: '#EF6C00',
  botFontColor: '#fff',
  userBubbleColor: '#fff',
  userFontColor: '#4a4a4a',
}

const steps = [
  {
    id: '1',
    message: 'Olá Olá, sou o Me lembra bot!, qual seu nome?',
    trigger: '2',
  },
  {
    id: '2',
    user: true,
    trigger: '3',
  },
  {
    id: '3',
    message: 'Boa {previousValue}, muito legal seu nome! agora me fala seu número de celular!',
    trigger: '4',
  },
  {
    id: '4',
    user: true,
    validator: (value) => {
      if (!validator(value)) {
        return 'número inválido, confira novamente não esqueça do DD'
      }

      return true
    },
    trigger: '5',
  },
  {
    id: '5',
    message: 'Boa estamos quase no fim, me fala! o recado que você quer lembrar?',
    trigger: '6',
  },
  {
    id: '6',
    user: true,
    trigger: '7',
  },
  {
    id: '7',
    message: 'Quando você quer que eu te lembre disso (para me ajudar, escreva em minutos) ?',
    trigger: '8',
  },
  {
    id: '8',
    user: true,
    validator: (value) => {
      if (isNaN(value)) {
        return 'apenas números e em minutos!'
      }

      if (parseInt(value) > 20) {
        return 'ainda não consigo lembrar tudo isso!'
      }

      if (parseInt(value) <= 0) {
        return 'muito pouco tempo, aumente ai rsrs'
      }

      return true
    },
    trigger: '9',
  },
  {
    id: '9',
    message: 'Perfeito! Me lembra Bot! vai te lembrar disso no tempo correto!',
    end: true,
  },
]

async function sendToBackEnd({ steps, values }) {
  const [name, phoneNumber, message, minutes] = values

  try {
    const response = await axios.post(`/reminder-me`, {
      name: name,
      phone_number: phoneNumber.replace(/\D/g, ""),
      message: message,
      reminder_in_minutes: parseInt(minutes),
    })

    console.log(response);
  } catch (error) {
    console.error(error);
  }
}


function App() {
  return (
    <ThemeProvider theme={theme}>
      <ChatBot handleEnd={sendToBackEnd} placeholder="Digite uma mensagem..." headerTitle="Me lembra Bot!" steps={steps} userAvatar="https://logospng.org/download/cruzeiro-do-sul/logo-cruzeiro-do-sul-estrela-256.png"/>
    </ThemeProvider>
  );
}

export default App;
