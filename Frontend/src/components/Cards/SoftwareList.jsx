import React from "react";
import { List, ListItem, ListItemText, Paper, Typography } from "@mui/material";

const SoftwareList = ({ software }) => {
  if (!software || software.length === 0) {
    return <Typography>No installed software found</Typography>;
  }

  return (
    <Paper sx={{ maxHeight: 600, overflow: "auto", padding: 2 }}>
      <Typography variant="h6" gutterBottom>
        Installed Software
      </Typography>
      <List>
        {software.map((item, index) => (
          <ListItem key={index} divider>
            <ListItemText primary={item} />
          </ListItem>
        ))}
      </List>
    </Paper>
  );
};

export default SoftwareList;
