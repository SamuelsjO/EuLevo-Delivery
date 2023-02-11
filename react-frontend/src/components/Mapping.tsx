import { Grid } from "@material-ui/core";
import Button from "@material-ui/core/Button";
import MenuItem from "@material-ui/core/MenuItem";
import Select from "@material-ui/core/Select";
import { useEffect, useState } from "react";
import { Route } from "../util/model";

//const API_URL = process.env.REACT_APP_API_URL;

type Props = {};
export const Mapping = (props: Props) => {
  const [routes, setRoutes] = useState<Route[]>([]);
  const [routesIdSelected, setRoutesIdselected] = useState<String>("");

  useEffect(() => {
    fetch(`http://localhost:3000/routes`)
      .then((data) => data.json())
      .then((data) => setRoutes(data));
  }, []);
  return (
    <Grid container>
      <Grid item xs={12} sm={3}>
        <form>
          <Select
            fullWidth
            displayEmpty
            value={routesIdSelected}
            onChange={(event) => setRoutesIdselected(event.target.value + "")}
          >
            <MenuItem value="">
              <em>Selecione uma corrida</em>
            </MenuItem>
            {routes.map((route, key) => (
              <MenuItem key={key} value={route.id}>
                {route.title}
              </MenuItem>
            ))}
          </Select>
          <Button type="submit" color="primary" variant="contained">
            Iniciar uma corrida
          </Button>
        </form>
      </Grid>
      <Grid item xs={12} sm={9}>
        <div id="map"></div>
      </Grid>
    </Grid>
  );
};
