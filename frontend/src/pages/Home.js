import {Button, CardMedia, Grid, Input, Paper, Slider} from "@mui/material"
import {Box} from "@mui/system"
import {useEffect, useState} from "react"

const Card = (props) => {
  const country = props.country
  return (
    <Grid item xs={3}>
      <Paper>
        <Box display="flex" flexDirection="column">
          {country.Name}
          <img 
            fullWidth
            height={100}
            src={country.Flag} alt=""/>
        </Box>
      </Paper>
    </Grid>
  )
}

const Population = (props) => {
  const value = props.value
  const setValue = props.setValue
  const handleChange = (event, newValue) => {
    setValue(newValue)
  }
  const handleInputChange = (event) => {
    console.log(event.target.value)
    setValue(event.target.value === '' ? '' : Number(event.target.value))
  }

  return (
    <Grid container sx={{ width: "50vw" }} spacing={2}>
      <Input 
        value={value[0]}
        size="small"
        onChange={handleInputChange}
        />
      <Slider 
        getAriaLabel={() => "Population Range"}
        valueLabelDisplay="auto"
        value={value}
        onChange={handleChange}
        min={100}
        max={2000000000}
        />
      <Input 
        value={value[1]}
        size="small"
        onChange={handleInputChange}
        />
    </Grid>
  )
}

function Home() {

  const [value, setValue] = useState([100, 2000000000])
  const [countries, setCountries] = useState([])
  const [isLoading, setIsLoading] = useState(true)

  const handleSubmit = () => {
    const url = `http://localhost:8080/filter?min=${value[0]}&max=${value[1]}`
    fetch(url, {
      method: 'GET',
      mode: 'cors',
    })
    .then(response => response.json())
    .then(data => {
        console.log(data)
        // setCountries(data)
      })
    .catch((error) => {
        console.error('Error:', error)
      })
  }

  useEffect(() => {
    fetch('http://localhost:8080/', {
      method: 'GET',
      mode: 'cors',
    })
      .then(response => response.json())
      .then(data => {
        setCountries(data)
      })
      .catch((error) => {
        console.error('Error:', error)
      })
  }, [])

  return (
    <Grid container spacing={2} p={4}>
      <Box p={2} m={2} fullWidth>
        <Population value={value} setValue={setValue}/>
      </Box>
      <Grid container item xs={4}>
        <Button onClick={handleSubmit}>hello</Button>
      </Grid>
      <Grid container item xs={10} spacing={2}>
        {countries.map((c, i) => <Card country={c} key={i} />)}
      </Grid>
    </Grid>
  )
}

export default Home
