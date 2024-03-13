package database

import (
	"context"
	"crud/app/models"
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// mock for now

func (p *Postgres) GetNodeParent(node models.GraphNode) (models.GraphNode, error) {
	table_name, _ := GetTable(node)

	// if err != nil needed ^

	parent, err := p.DB.Query(context.Background(), `
	SELECT id, type, name, parent
	FROM $1
	WHERE id=$2
	
	`, table_name, node.Parent)

	if err != nil {
		p.Log.Error("Could not get parent pgx.Rows",
			zap.Error(err))
	}

	parentNode, err := pgx.CollectOneRow[models.GraphNode](parent, pgx.RowToStructByName[models.GraphNode])

	if err != nil {
		p.Log.Error("Error while getting parent node. The node probably is ROOT",
			zap.Error(err))
		return models.GraphNode{}, nil
	}

	return parentNode, nil
}

func GetTable(node models.GraphNode) (string, error) {
	var (
		table_name string
		err        error
	)
	switch node_type := node.Type; node_type {
	case "location":
		table_name = "location_ids"
		err = nil

	case "microcategory":
		table_name = "microcategory_ids"
		err = nil

	default:
		table_name = ""
		err = nil

	}

	return table_name, err
}

func ConvertRowsToGraphNode(rows pgx.Rows) (models.GraphNode, error) {
	var (
		node models.GraphNode
	)

	defer rows.Close()

	fmt.Println(rows)

	return node, nil
}
