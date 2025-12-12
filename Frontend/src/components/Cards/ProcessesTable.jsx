import React from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
} from "@mui/material";

const ProcessesTable = ({ processes }) => {
  if (!processes || processes.length === 0) {
    return <p>No process data available</p>;
  }

  return (
    <TableContainer component={Paper} sx={{ maxHeight: 600 }}>
      <Table stickyHeader>
        <TableHead>
          <TableRow>
            <TableCell>PID</TableCell>
            <TableCell>Name</TableCell>
            <TableCell>CPU (%)</TableCell>
            <TableCell>Memory (%)</TableCell>
          </TableRow>
        </TableHead>

        <TableBody>
          {processes.map((proc) => (
            <TableRow key={proc.pid}>
              <TableCell>{proc.pid}</TableCell>
              <TableCell>{proc.name}</TableCell>
              <TableCell>{(proc.cpu_percent ?? 0).toFixed(2)}</TableCell>
              <TableCell>{(proc.mem_percent ?? 0).toFixed(2)}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default ProcessesTable;
