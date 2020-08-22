package constanta

const (
	// GETALLRESERVATION = `SELECT mr.id_room , mr.name_room,mr.floor_location,mr.price,mrs.id_status_room,mrs.name_status_room
	GETALLRESERVATION = `SELECT mr.id_room , mr.name_room,mr.floor_location,mr.price,mrs.id_status_room,mrs.name_status_room
							  FROM m_room mr JOIN m_status_room mrs ON mr.status_room = mrs.id_status_room  ORDER BY mr.id_room ASC;`
	CREATERESERVATION = `INSERT INTO m_room VALUES (?, ?, ?, ?, ?);`
)
