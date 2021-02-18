package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
	"github.com/lukasbischof/tech_radar/models"
	"net/http"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Technology)
// DB Table: Plural (technologies)
// Resource: Plural (Technologies)
// Path: Plural (/technologies)
// View Template Folder: Plural (/templates/technologies/)

// TechnologiesResource is the resource for the Technology model
type TechnologiesResource struct {
	buffalo.Resource
}

// List gets all Technologies. This function is mapped to the path
// GET /technologies
func (v TechnologiesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	technologies := &models.Technologies{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Technologies from the DB
	if err := q.All(technologies); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("technologies", technologies)
		return c.Render(http.StatusOK, r.HTML("/technologies/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(technologies))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(technologies))
	}).Respond(c)
}

// Show gets the data for one Technology. This function is mapped to
// the path GET /technologies/{technology_id}
func (v TechnologiesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Technology
	technology := &models.Technology{}

	// To find the Technology the parameter technology_id is used.
	if err := tx.Find(technology, c.Param("technology_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("technology", technology)

		return c.Render(http.StatusOK, r.HTML("/technologies/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(technology))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(technology))
	}).Respond(c)
}

// New renders the form for creating a new Technology.
// This function is mapped to the path GET /technologies/new
func (v TechnologiesResource) New(c buffalo.Context) error {
	c.Set("technology", &models.Technology{})

	return c.Render(http.StatusOK, r.HTML("/technologies/new.plush.html"))
}

// Create adds a Technology to the DB. This function is mapped to the
// path POST /technologies
func (v TechnologiesResource) Create(c buffalo.Context) error {
	// Allocate an empty Technology
	technology := &models.Technology{}

	// Bind technology to the html form elements
	if err := c.Bind(technology); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(technology)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("technology", technology)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/technologies/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "technology.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/technologies/%v", technology.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(technology))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(technology))
	}).Respond(c)
}

// Edit renders a edit form for a Technology. This function is
// mapped to the path GET /technologies/{technology_id}/edit
func (v TechnologiesResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Technology
	technology := &models.Technology{}

	if err := tx.Find(technology, c.Param("technology_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("technology", technology)
	return c.Render(http.StatusOK, r.HTML("/technologies/edit.plush.html"))
}

// Update changes a Technology in the DB. This function is mapped to
// the path PUT /technologies/{technology_id}
func (v TechnologiesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Technology
	technology := &models.Technology{}

	if err := tx.Find(technology, c.Param("technology_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Technology to the html form elements
	if err := c.Bind(technology); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(technology)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("technology", technology)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/technologies/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "technology.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/technologies/%v", technology.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(technology))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(technology))
	}).Respond(c)
}

// Destroy deletes a Technology from the DB. This function is mapped
// to the path DELETE /technologies/{technology_id}
func (v TechnologiesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Technology
	technology := &models.Technology{}

	// To find the Technology the parameter technology_id is used.
	if err := tx.Find(technology, c.Param("technology_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(technology); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "technology.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/technologies")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(technology))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(technology))
	}).Respond(c)
}
