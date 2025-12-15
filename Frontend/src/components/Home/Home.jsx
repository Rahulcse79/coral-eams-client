import React, { useEffect, useState } from "react";
import {
  Card,
  CardContent,
  Typography,
  Table,
  TableHead,
  TableBody,
  TableRow,
  TableCell,
  IconButton,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import VisibilityIcon from "@mui/icons-material/Visibility";
import "./Home.css";

export default function Home() {
  const [assetsList, setAssetsList] = useState([]);

  const navigate = useNavigate();

  useEffect(() => {
    const dummyAssets = [
      { id: 1, mac_address: "f0:79:60:04:f3:16", device: "MacBook Air" },
      { id: 2, mac_address: "00:07:32:bf:25:7e", device: "Linux server" },
      { id: 3, mac_address: "00:ff:d3:97:d9:62", device: "Windows laptop" },
    ];
    setAssetsList(dummyAssets);
  }, []);

  return (
    <div className="v-eams-home-page">
      <Card className="v-asset-card" variant="outlined">
        <CardContent>
          <Typography variant="h6" className="v-card-title">
            Assets List
          </Typography>

          <Table className="v-asset-table">
            <TableHead>
              <TableRow>
                <TableCell>
                  <b>ID</b>
                </TableCell>
                <TableCell>
                  <b>Device</b>
                </TableCell>
                <TableCell>
                  <b>MAC Address</b>
                </TableCell>
                <TableCell align="center">
                  <b>Action</b>
                </TableCell>
              </TableRow>
            </TableHead>

            <TableBody>
              {assetsList.map((item) => (
                <TableRow key={item.id} className="v-table-row">
                  <TableCell>{item.id}</TableCell>
                  <TableCell>{item.device}</TableCell>
                  <TableCell>{item.mac_address}</TableCell>
                  <TableCell align="center">
                    <IconButton
                      color="primary"
                      onClick={() => navigate(`/dashboard/${item.mac_address}`)}
                    >
                      <VisibilityIcon />
                    </IconButton>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </CardContent>
      </Card>
    </div>
  );
}
